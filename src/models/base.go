package models

import (
	"time"

	"gorm.io/gorm"
)

/**
 * @model Base
 * @description This is injected into other models to provide common functionality.
 */

type Base struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
