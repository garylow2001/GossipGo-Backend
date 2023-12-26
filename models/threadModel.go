package models

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	ID       int       `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Body     string    `json:"body,omitempty"`
	AuthorID int       `gorm:"foreignKey:AuthorID" json:"author_id,omitempty"`
	Author   User      `json:"author,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}
