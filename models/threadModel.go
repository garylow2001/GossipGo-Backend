package models

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	Title      string       `gorm:"not null" json:"title,omitempty"`
	Body       string       `json:"body,omitempty"`
	AuthorID   uint         `gorm:"not null" json:"author_id,omitempty"`
	Author     User         `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Comments   []Comment    `gorm:"foreignKey:ThreadID" json:"comments,omitempty"`
	Category   string       `json:"category,omitempty"`
	Likes      []ThreadLike `gorm:"foreignKey:ThreadID" json:"likes,omitempty"`
	LikesCount int          `json:"likes_count"`
}

var ValidCategories = []string{"Academic", "News", "Technology", "Entertainment", "Hot Takes"}
