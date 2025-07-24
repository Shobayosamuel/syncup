package service

import (

	"github.com/Shobayosamuel/syncup/services/user-service/internal/repository"
)

type userProfileService struct {
	userProfileRepo repository.UserProfileRepository
}

func NewUserProfileService(userProfileRepo repository.UserProfileRepository)