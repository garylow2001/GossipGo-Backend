package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id,omitempty"`
	Username string `gorm:"unique" json:"username,omitempty"`
	Auth     Auth   `gorm:"foreignKey:UserID" json:"auth,omitempty"`
	// Threads  []Thread  `json:"threads,omitempty"`
	// Comments []Comment `json:"comments,omitempty"`
}
