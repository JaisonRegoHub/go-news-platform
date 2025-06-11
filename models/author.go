package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name        string    `json:"name"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
	Articles    []Article `json:"-"`
}
