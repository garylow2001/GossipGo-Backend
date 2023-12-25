package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	UserID   uint   `gorm:"primaryKey" json:"user_id,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}
