// models/user.go
package models

import (
	"gorm.io/gorm"
)

// User 是用户模型
type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}
