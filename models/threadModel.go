package models

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"title,omitempty"`
	Body     string    `json:"body,omitempty"`
	AuthorID uint      `gorm:"not null" json:"author_id,omitempty"`
	Author   User      `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Comments []Comment `gorm:"foreignKey:ThreadID" json:"comments,omitempty"`
}
