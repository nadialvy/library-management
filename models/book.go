package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"size:255;not null"`
	Author      string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
}
