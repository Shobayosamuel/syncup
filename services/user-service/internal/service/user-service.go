package service

import (
	"github.com/Shobayosamuel/syncup/services/user-service/internal/repository"
	"github.com/Shobayosamuel/syncup/shared/models"
	"github.com/Shobayosamuel/syncup/shared/utils"
	"github.com/google/uuid"
)

type userProfileService struct {
	userProfileRepo repository.UserProfileRepository
}

type UserProfileService interface {
	GetProfile(userID uuid.UUID) (*models.UserProfile, error)
	UpdateProfile(userID uuid.UUID, updateData utils.UpdateProfileRequest) (*models.UserProfile, error)
	AddPhoto(userID uuid.UUID, photoURL string) error
	DeletePhoto(userID uuid.UUID, photoID string) error
	UpdatePreferences(userID uuid.UUID, preferences utils.UpdatePreferencesRequest) (*models.UserPreference, error)
	GetRecommendations(userID uuid.UUID) ([]models.UserProfile, error)
	UpdateLocation(userID uuid.UUID, lat, lng float64) error
}

func NewUserProfileService(userProfileRepo repository.UserProfileRepository) UserProfileService {
	return &userProfileService{
		userProfileRepo: userProfileRepo,
	}
}

func (s *userProfileService) GetProfile(userID uuid.UUID) (*models.UserProfile, error) {
	return s.GetProfile(userID)
}

func (s *userProfileService) UpdateProfile(userID uuid.UUID, updateData utils.UpdateProfileRequest) (*models.UserProfile, error) {
	profile, err := s.userProfileRepo.GetProfile(userID)
	if err != nil {
		// create a new profile if itis non-existent
		profile = &models.UserProfile{
			UserID: userID,
		}
	}
	if updateData.Name != "" {
		profile.Name = updateData.Name
	}
	if updateData.Bio != "" {
		profile.Bio = updateData.Bio
	}
	if updateData.Age > 0 {
		profile.Age = updateData.Age
	}
	if updateData.Gender != "" {
		profile.Gender = updateData.Gender
	}
	if len(updateData.Interests) > 0 {
		profile.Interests = updateData.Interests
	}
	if profile.ID == uuid.Nil {
		err = s.userProfileRepo.CreateProfile(profile)
	} else {
		err = s.userProfileRepo.UpdateProfile(profile)
	}
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (s *userProfileService) AddPhoto(userID uuid.UUID, photoURL string) error {
	profile, err := s.userProfileRepo.GetProfile(userID)
	if err != nil {
		return err
	}
	if len(proflle.Photos)
}