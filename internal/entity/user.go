package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/fabioods/go-orders/internal/errorcode"
	errorformated "github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/trace"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id" validate:"required,uuid4" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" validate:"required" gorm:"type:varchar(255); not null"`
	AvatarURL string    `json:"avatar_url"`
	Email     string    `json:"email" validate:"required,email" gorm:"type:varchar(255); not null; uniqueIndex"`
	Password  string    `json:"-" validate:"required" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at"  gorm:"autoUpdateTime"`
}

func NewUser() *User {
	return &User{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) DefineAvatar(avatar string) {
	if len(avatar) == 0 {
		return
	}
	u.AvatarURL = avatar
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		var errMsgs []string
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf("Field '%s' invalid: %s", err.Field(), err.Tag()))
		}
		return fmt.Errorf("%s", strings.Join(errMsgs, "; "))
	}
	return nil
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return errorformated.BadRequestError(trace.GetTrace(), errorcode.ErrorUserValidate, "Password is required")
	}
	hash, err := u.hashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ef := errorformated.UnexpectedError(trace.GetTrace(), errorcode.ErrorUserBcryptError, "%s", err.Error())
		return "", ef
	}

	return string(bytes), nil
}
