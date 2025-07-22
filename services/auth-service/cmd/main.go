package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Shobayosamuel/syncup/services/auth-service/internal/config"
	"github.com/Shobayosamuel/syncup/shared/models"
	"github.com/Shobayosamuel/syncup/services/auth-service/internal/repository"
	"github.com/Shobayosamuel/syncup/services/auth-service/internal/service/auth_service"
)


func main() {
	// Load config
	cfg := config.Load()

	// Setup database
	db := setupDatabase(cfg)

	// Auto migrate
	db.AutoMigrate(&models.User{})

	// Setup repositories
	userRepo := repository.NewUserRepository(db)

	// Setup services
	authService := auth_service.NewUserService(userRepo)


	// Setup handlers
	authHandler := auth.NewHandler(authService)

	// Setup router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Public routes
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/refresh", authHandler.RefreshToken)
	}


	// Protected routes
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AuthMiddleware(authService))
	{
		// Auth routes
		apiGroup.GET("/profile", authHandler.GetProfile)

	}

	// Start server
	log.Printf("Server starting on :%s", cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}

func setupDatabase(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}
