package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string        `gorm:"unique" json:"username,omitempty"`
	Threads      []Thread      `gorm:"foreignKey:AuthorID" json:"threads,omitempty"`
	Comments     []Comment     `gorm:"foreignKey:UserID" json:"comments,omitempty"`
	ThreadLikes  []ThreadLike  `gorm:"foreignKey:UserID" json:"thread_likes,omitempty"`
	CommentLikes []CommentLike `gorm:"foreignKey:UserID" json:"comment_likes,omitempty"`
}
