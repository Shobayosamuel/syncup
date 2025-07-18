package auth

type RegisterRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min-6"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min-6"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}