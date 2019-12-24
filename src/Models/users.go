package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User object...
type User struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Idregimen  uint      `json:"regimen"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

//BeforeCreate User set when was created
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Created_at", time.Now())
	return nil
}

//BeforeUpdate set when its updated a register
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Updated_at", time.Now())
	return nil
}
