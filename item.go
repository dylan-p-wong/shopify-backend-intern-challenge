package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Item struct {
	ID           uuid.UUID    `gorm:"type:uuid" json:"id"`
	Title        string       `json:"title" binding:"required"`
	Description  string       `json:"description"`
  Count        int          `json:"count" binding:"required"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

func (queue *Item) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	queue.ID = id
	return err
}
