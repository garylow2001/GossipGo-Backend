package models

import "gorm.io/gorm"

type CommentLike struct {
	gorm.Model
	UserID    uint `gorm:"not null" json:"user_id,omitempty"`
	CommentID uint `gorm:"not null" json:"comment_id,omitempty"`
}
