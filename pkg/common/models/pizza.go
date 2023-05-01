package models

import (
	"github.com/google/uuid"
	"time"
)

type Pizza struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
