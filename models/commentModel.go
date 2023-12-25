package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID     int    `json:"id,omitempty"`
	Body   string `json:"body,omitempty"`
	Author User   `json:"author,omitempty"`
	Thread Thread `json:"thread,omitempty"`
}
