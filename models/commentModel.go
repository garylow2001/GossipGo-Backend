package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body      string `json:"body,omitempty"`
	UserID    uint   `gorm:"not null" json:"author_id,omitempty"`
	Author    User   `gorm:"foreignKey:UserID" json:"author,omitempty"`
	ThreadID  uint   `gorm:"not null" json:"thread_id,omitempty"`
	Thread    Thread `gorm:"foreignKey:ThreadID" json:"thread,omitempty"`
	CommentID uint   `gorm:"not null" json:"comment_id,omitempty"`
}
