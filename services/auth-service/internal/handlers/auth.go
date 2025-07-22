package handlers

import (
	"net/http"
	"github.com/Shobayosamuel/syncup/shared/utils"
	"github.com/gin-gonic/gin"
	"github.com/Shobayosamuel/syncup/shared/models"
	"github.com/Shobayosamuel/syncup/shared/interfaces"
)

type Handler struct {
	service interfaces.UserService
}

func NewHandler(service interfaces.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c *gin.Context) {
	var req utils.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, _ := h.service.GetUserFromToken(tokens.AccessToken)
	userResponse := utils.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		IsActive: user.IsActive,
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"tokens":  tokens,
		"user":    userResponse,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req utils.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	user, _ := h.service.GetUserFromToken(tokens.AccessToken)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"tokens":  tokens,
		"user": utils.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			IsActive: user.IsActive,
		},
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var req utils.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token refreshed successfully",
		"tokens":  tokens,
	})
}

func (h *Handler) GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	userModel := user.(*models.User)
	userResponse := utils.UserResponse{
		ID:       userModel.ID,
		Email:    userModel.Email,
		IsActive: userModel.IsActive,
	}

	c.JSON(http.StatusOK, gin.H{"user": userResponse})
}
