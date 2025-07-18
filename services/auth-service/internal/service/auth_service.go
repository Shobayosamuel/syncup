package service

import (
	"errors"

	"github.com/Shobayosamuel/syncup/services/auth-service/internal/repository"
	"github.com/Shobayosamuel/syncup/shared/auth"
	"github.com/Shobayosamuel/syncup/shared/models"
)

type UserService interface {
	Register(req auth.RegisterRequest) (*auth.TokenResponse, error)
	Login(req auth.LoginRequest) (*auth.TokenResponse, error)
	RefreshToken(req string) (*auth.TokenResponse, error)
	GetUserFromToken(req string) (*models.User, error)
	GenerateTokens(user models.User) (*auth.TokenResponse, error)
}
type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(req auth.RegisterRequest) (*auth.TokenResponse, error) {
	if _, err := s.userRepo.GetByEmail(req.Email); err == nil {
		return nil, errors.New("email already exists")
	}
	hashed_password, err := auth.HashPassword(req.Password)
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

func (s *userService) Login(req auth.LoginRequest) (*auth.TokenResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email already exists")
	}
	match := auth.CheckPassword(req.Password, user.Password)
	if !match {
		return nil, errors.New("invalid credentials")
	}
	return s.GenerateTokens(*user)
}

func (s *userService) RefreshToken(req string) (*auth.TokenResponse, error) {
	// validate the token
	claims, err := auth.ValidateRefreshToken(req)
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

func (s *userService) GenerateTokens(user models.User) (*auth.TokenResponse, error) {
	accessToken, err := auth.GenerateAccessToken(user)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := auth.GenerateRefreshToken(user)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &auth.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(auth.GetAccessTokenTTL().Seconds()),
	}, nil
}

func (s *userService) GetUserFromToken(req string) (*models.User, error) {
	claims, err := auth.ValidateAccessToken(req)
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