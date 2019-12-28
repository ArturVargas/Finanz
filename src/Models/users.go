package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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

// BeforeSave we need hash password
func (user *User) BeforeSave() error {
	hashedPwd, err := HashingPwd(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPwd)
	return nil
}

// HashingPwd for Hashing Password
func HashingPwd(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPwd is for find user password in db
func VerifyPwd(hashedPwd, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
}
