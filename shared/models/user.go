package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"uniqueIndex; not null"`
	Password string `json:"-" gorm:"not null"`
	IsActive bool `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserProfile struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID      `json:"user_id" gorm:"type:uuid;unique;not null"`
	Name      string         `json:"name" gorm:"size:100;not null"`
	Bio       string         `json:"bio"`
	Age       int            `json:"age"`
	Gender    string         `json:"gender" gorm:"size:20"`
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	Photos    pq.StringArray `json:"photos" gorm:"type:text[]"`
	Interests pq.StringArray `json:"interests" gorm:"type:text[]"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type UserPreference struct {
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID          uuid.UUID `json:"user_id" gorm:"type:uuid;unique;not null"`
	MinAge          int       `json:"min_age" gorm:"default:18"`
	MaxAge          int       `json:"max_age" gorm:"default:100"`
	MaxDistance     int       `json:"max_distance" gorm:"default:50"`
	PreferredGender string    `json:"preferred_gender" gorm:"size:20"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}