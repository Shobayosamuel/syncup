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

func NewUserProfileService(userProfileRepo repository.UserProfileRepository)