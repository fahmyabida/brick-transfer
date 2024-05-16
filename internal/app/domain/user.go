package domain

import (
	"time"
)

type Users struct {
	ID        string     `json:"id" gorm:"id;primary_key"`
	Name      string     `json:"name" gorm:"name"`
	Type      string     `json:"type" gorm:"type"`
	Status    string     `json:"status" gorm:"status"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamptz;NOT NULL;default:null"`
}
