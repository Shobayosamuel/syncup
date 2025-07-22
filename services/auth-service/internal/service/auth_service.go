package service

import (
	"errors"

	"github.com/Shobayosamuel/syncup/services/auth-service/internal/repository"
	"github.com/Shobayosamuel/syncup/shared/utils"
	"github.com/Shobayosamuel/syncup/shared/models"
	"github.com/Shobayosamuel/syncup/shared/interfaces"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) interfaces.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(req utils.RegisterRequest) (*utils.TokenResponse, error) {
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("email already exists")
	}
	hashed_password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	// create the user
	user := &models.User{
		Email: req.Email,
		Password: hashed_password,
		IsActive: true,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create a new user")
	}
	return s.GenerateTokens(*user)
}

func (s *userService) Login(req utils.LoginRequest) (*utils.TokenResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email already exists")
	}
	match := utils.CheckPassword(req.Password, user.Password)
	if !match {
		return nil, errors.New("invalid credentials")
	}
	return s.GenerateTokens(*user)
}

func (s *userService) RefreshToken(req string) (*utils.TokenResponse, error) {
	// validate the token
	claims, err := utils.ValidateRefreshToken(req)
	if err != nil {
		return nil, err
	}
	// get user from claims
	user, err := s.userRepo.GetByID(claims.UserID)
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user is inactive")
	}
	return s.GenerateTokens(*user)
}

func (s *userService) GenerateTokens(user models.User) (*utils.TokenResponse, error) {
	accessToken, err := utils.GenerateAccessToken(user)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &utils.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(utils.GetAccessTokenTTL().Seconds()),
	}, nil
}

func (s *userService) GetUserFromToken(req string) (*models.User, error) {
	claims, err := utils.ValidateAccessToken(req)
	if err != nil {
		return nil, err
	}
	// get user from claims
	user, err := s.userRepo.GetByID(claims.UserID)
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, errors.New("user is inactive")
	}
	return user, nil
}