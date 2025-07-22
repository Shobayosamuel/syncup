package interfaces

import (
	"github.com/Shobayosamuel/syncup/shared/models"
	"github.com/Shobayosamuel/syncup/shared/utils"
)

type UserService interface {
	Register(req utils.RegisterRequest) (*utils.TokenResponse, error)
	Login(req utils.LoginRequest) (*utils.TokenResponse, error)
	RefreshToken(req string) (*utils.TokenResponse, error)
	GetUserFromToken(req string) (*models.User, error)
	GenerateTokens(user models.User) (*utils.TokenResponse, error)
}