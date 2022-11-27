package model

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy uuid.UUID  `json:"created_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy uuid.UUID  `json:"updated_by"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	DeletedBy uuid.UUID  `json:"deleted_by"`
}
