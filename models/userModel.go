package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID       uint   `json:"id,omitempty"`
	Username string `gorm:"unique" json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	// Threads  []Thread  `json:"threads,omitempty"`
	// Comments []Comment `json:"comments,omitempty"`
}
