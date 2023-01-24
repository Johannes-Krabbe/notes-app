package models

import "gorm.io/gorm"

type Note struct {
    gorm.Model  // adds ID, created_at etc.
    Title       string `json:"title"`
    Content 	string `json:"content"`
}