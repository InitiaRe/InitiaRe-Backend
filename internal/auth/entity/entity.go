package entity

import (
	"time"

	"InitiaRe-website/constant"
	"InitiaRe-website/internal/auth/models"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Id        int       `gorm:"primarykey;column:id" json:"id"`
	UserId    uuid.UUID `gorm:"column:user_id" json:"user_id"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	Email     string    `gorm:"column:email" json:"email,omitempty"`
	Password  string    `gorm:"column:password" json:"password,omitempty"`
	School    string    `gorm:"column:school" json:"school"`
	Gender    string    `gorm:"column:gender" json:"gender,omitempty"`
	Birthday  time.Time `gorm:"column:birthday;default:(-)" json:"birthday,omitempty"`
	CreatedAt time.Time `gorm:"autoCreatetime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:(-)" json:"updated_at,omitempty"`
	LoginDate time.Time `gorm:"column:login_date;default:(-)" json:"login_date,omitempty"`

	// Custom fields
	Status int `gorm:"->;-:migration" json:"status,omitempty"`
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
	if u.Status == 0 {
		u.Status = constant.USER_STATUS_ACTIVE
	}
	u.newUUID()
	u.HashPassword()
}
