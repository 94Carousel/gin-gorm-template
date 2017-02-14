package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	BaseModel
	Username            string
	Nickname            string `sql:"index"`
	Mobile              string `sql:"index"`
	Email               string
	Avatar              string
	Password            string `gorm:"-"`
	EncryptedPassword   string
	AuthToken           string
	SignInCount         int
	LastSignInAt        time.Time
	ResetPasswordToken  string
	ResetPasswordSentAt *time.Time
}

// ValidPassword valid password validity; return bool
func (user User) ValidPassword(pwd string) bool {
	password := []byte(pwd)
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), password)
	if err != nil {
		return false
	}
	return true
}

// GeneratePassword generate a encrypt password
func (user *User) GeneratePassword() {
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.EncryptedPassword = string(hashedPassword)
}

// SignIn user sign in action
func (user *User) SignIn() bool {
	if !user.ValidPassword(user.Password) {
		return false
	}
	DB.Model(&user).Update("last_sign_in_at", time.Now())
	return true
}
