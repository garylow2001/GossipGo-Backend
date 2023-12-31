package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"unique" json:"username,omitempty"`
	Threads  []Thread  `gorm:"foreignKey:AuthorID" json:"threads,omitempty"`
	Comments []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`
	// Auth     Auth   `gorm:"foreignKey:UserID" json:"auth,omitempty"`
}
