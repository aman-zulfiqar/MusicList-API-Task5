package models

import "github.com/google/uuid"

type Song struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Genre  string    `json:"genre"`
}
