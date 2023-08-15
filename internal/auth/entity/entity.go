package entity

import (
	"time"

	"github.com/Ho-Minh/InitiaRe-website/internal/auth/models"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
	Birthday  time.Time `gorm:"column:birthday;default:(-)" json:"birthday,omitempty" redis:"birthday"`
	CreatedAt time.Time `gorm:"autoCreatetime" json:"created_at,omitempty" redis:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty" redis:"updated_at"`
	LoginDate time.Time `gorm:"column:login_date;default:(-)" json:"login_date,omitempty" redis:"login_date"`
}

func (u *User) TableName() string {
	return "initiaRe_user"
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

// Sanitize user password
func (u *User) SanitizePassword() {
	u.Password = ""
}

// Generate new user id
func (u *User) newUUID() {
	u.UserId = uuid.New()
}

func (u *User) Export() *models.Response {
	obj := &models.Response{}
	copier.Copy(obj, u) //nolint
	return obj
}

func (u *User) ParseFromSaveRequest(req *models.SaveRequest) {
	copier.Copy(u, req) //nolint
	u.newUUID()
	u.HashPassword()
}
