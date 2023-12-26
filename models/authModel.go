package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	User     User   `gorm:"foreignKey:UserID"`
	UserID   uint   `gorm:"primaryKey" json:"user_id,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}
