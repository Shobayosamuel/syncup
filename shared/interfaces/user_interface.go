package interfaces

// import (
// 	"github.com/Shobayosamuel/syncup/shared/utils"
// 	"github.com/google/uuid"
// 	"github.com/Shobayosamuel/syncup/services/user-service/repository"
// )

// type UserProfileService interface {
// 	GetProfile(userID uuid.UUID) (*repository.UserProfile, error)
// 	UpdateProfile(userID uuid.UUID, updateData utils.UpdateProfileRequest) (*repository.UserProfile, error)
// 	AddPhoto(userID uuid.UUID, photoURL string) error
// 	DeletePhoto(userID uuid.UUID, photoID string) error
// 	UpdatePreferences(userID uuid.UUID, preferences utils.UpdatePreferencesRequest) (*repository.UserPreference, error)
// 	GetRecommendations(userID uuid.UUID) ([]repository.UserProfile, error)
// 	UpdateLocation(userID uuid.UUID, lat, lng float64) error
// }