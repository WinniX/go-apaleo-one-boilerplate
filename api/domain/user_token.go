package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	UserUID      uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	AccountCode  string    `gorm:"not null;size:10"`
	AccessToken  string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	Expiry       time.Time `gorm:"type:timestamptz;not null"`
}
