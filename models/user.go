package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:enum('user', 'admin', 'librarian');default:'user';not null"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	validRoles := map[string]bool{
		"user":      true,
		"admin":     true,
		"librarian": true,
	}

	if !validRoles[u.Role] {
		return errors.New("invalid role")
	}

	return nil
}
