package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Id        int       `gorm:"primarykey;column:id" json:"id" redis:"id"`
	UserId    uuid.UUID `gorm:"column:user_id" json:"user_id" redis:"user_id"`
	FirstName string    `gorm:"column:first_name" json:"first_name" redis:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name" redis:"last_name"`
	Email     string    `gorm:"column:email" json:"email,omitempty" redis:"email"`
	Password  string    `gorm:"column:password" json:"password,omitempty" redis:"password"`
	School    string    `gorm:"column:school" json:"school" redis:"school"`
	Gender    string    `gorm:"column:gender" json:"gender,omitempty" redis:"gender"`
	Birthday  time.Time `gorm:"column:birthday" json:"birthday,omitempty" redis:"birthday"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at,omitempty" redis:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at,omitempty" redis:"updated_at"`
	LoginDate time.Time `gorm:"column:login_date" json:"login_date,omitempty" redis:"login_date"`
}

// Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Compare user password with bcrypt
func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

// Sanitize user password
func (u *User) SanitizePassword() {
	u.Password = ""
}

// Generate new user id
func (u *User) NewUUID() {
	u.UserId = uuid.New()
}

// Parse from request body
func (u *User) Parse() *User {
	u.NewUUID()
	u.HashPassword()
	return &User{
		UserId:    u.UserId,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Birthday:  u.Birthday,
		Gender:      u.Gender,
	}
}

func (u *User) TableName() string {
	return "users"
}

// User sign in response
type UserWithToken struct {
	User  *User  `json:"user,omitempty"`
	Token string `json:"token"`
}
