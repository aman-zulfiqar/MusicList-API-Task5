package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Song struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Genre  string    `json:"genre"`
	gorm.Model
}