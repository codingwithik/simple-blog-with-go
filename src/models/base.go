package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**
 * @model Base
 * @description This is injected into other models to provide common functionality.
 */

type Base struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.NewString()
	}

	return nil
}
