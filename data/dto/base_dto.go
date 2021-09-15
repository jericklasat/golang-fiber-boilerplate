package dto

import (
	"time"

	"gorm.io/gorm"
)

type BaseDto struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}