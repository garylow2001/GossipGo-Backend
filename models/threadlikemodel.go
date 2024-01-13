package models

import "gorm.io/gorm"

type ThreadLike struct {
	gorm.Model
	UserID   uint `gorm:"not null" json:"user_id,omitempty"`
	ThreadID uint `gorm:"not null" json:"thread_id,omitempty"`
}
