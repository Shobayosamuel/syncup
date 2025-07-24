package utils

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

type UpdateProfileRequest struct {
	Name      string   `json:"name"`
	Bio       string   `json:"bio"`
	Age       int      `json:"age"`
	Gender    string   `json:"gender"`
	Interests []string `json:"interests"`
}

type UpdatePreferencesRequest struct {
	MinAge          int    `json:"min_age"`
	MaxAge          int    `json:"max_age"`
	MaxDistance     int    `json:"max_distance"`
	PreferredGender string `json:"preferred_gender"`
}