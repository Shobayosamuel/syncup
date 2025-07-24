package repository

import (
	"errors"
	"sort"

	"github.com/Shobayosamuel/syncup/shared/models"
	"github.com/Shobayosamuel/syncup/shared/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserProfileRepository interface {
	GetProfile(userID uuid.UUID) (*models.UserProfile, error)
	CreateProfile(profile *models.UserProfile) error
	UpdateProfile(profile *models.UserProfile) error
	GetPreferences(userID uuid.UUID) (*models.UserPreference, error)
	UpdatePreferences(preferences *models.UserPreference) error
	GetRecommendations(userID uuid.UUID, limit int) ([]models.UserProfile, error)
}

type userProfileRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserProfileRepository {
	return &userProfileRepository{db: db}
}

func (r *userProfileRepository) GetProfile(userID uuid.UUID) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.Where("user_id == ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil

}

func (r *userProfileRepository) CreateProfile(profile *models.UserProfile) error {
	return r.db.Create(profile).Error
}

func (r *userProfileRepository) UpdateProfile(profile *models.UserProfile) error {
	return r.db.Save(profile).Error
}

func (r *userProfileRepository) GetPreferences(userID uuid.UUID) (*models.UserPreference, error) {
	var preferences models.UserPreference
	err := r.db.Where("user_id == ?", userID).First(&preferences).Error
	if err != nil {
		return nil, err
	}
	return &preferences, nil
}

func (r *userProfileRepository) UpdatePreferences(preferences *models.UserPreference) error {
	return r.db.Save(preferences).Error
}

func (r *userProfileRepository) GetRecommendations(userID uuid.UUID, limit int) ([]models.UserProfile, error) {
	var userPrefs models.UserPreference
	if err := r.db.Where("user_id = ?", userID).First(&userPrefs).Error; err != nil {
		userPrefs = models.UserPreference{MinAge: 18, MaxAge: 100, MaxDistance: 50}
	}

	var currentUser models.UserProfile
	if err := r.db.Where("user_id = ?", userID).First(&currentUser).Error; err != nil {
		return nil, err
	}

	if currentUser.Latitude == 0 || currentUser.Longitude == 0 {
		return nil, errors.New("user location is not set")
	}

	// base query
	query := r.db.Model(&models.UserProfile{}).
		Where("user_id != ? AND is_active = true", userID).
		Where("age BETWEEN ? AND ?", userPrefs.MinAge, userPrefs.MaxAge)

	// get gender preference
	if userPrefs.PreferredGender != "" {
		query = query.Where("gender = ?", userPrefs.PreferredGender)
	}

	// filter by distance (using Haversine, which is just a method to calculate distance cos of the unique shape of the earth )
	distanceSQL := utils.HaversineSQL(currentUser.Latitude, currentUser.Longitude)
	query = query.Select("*").Where(distanceSQL+" <= ?", userPrefs.MaxDistance)

	var candidates []models.UserProfile
	if err := query.Find(&candidates).Error; err != nil {
		return nil, err
	}

	// Sort by mutual interest count
	scored := []struct {
		Profile models.UserProfile
		Score   int
	}{}

	for _, profile := range candidates {
		score := utils.CountMutualInterests(currentUser.Interests, profile.Interests)
		scored = append(scored, struct {
			Profile models.UserProfile
			Score   int
		}{Profile: profile, Score: score})
	}

	// Sort by score descending
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	var recommendations []models.UserProfile
	for i, candidate := range scored {
		if i >= limit {
			break
		}
		recommendations = append(recommendations, candidate.Profile)
	}

	return recommendations, nil
}
