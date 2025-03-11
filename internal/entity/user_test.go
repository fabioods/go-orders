package entity

import (
	"testing"

	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/stretchr/testify/assert"
)

func TestUser_Valid(t *testing.T) {
	user := NewUser()
	user.Name = "Fabio"
	user.Email = "fabio@fabio.com"
	errPassword := user.SetPassword("fabio10")
	user.DefineAvatar("image1.jpg")

	assert.Equal(t, "image1.jpg", user.AvatarURL)
	assert.NotNil(t, user)
	assert.Equal(t, "fabio@fabio.com", user.Email)
	assert.Equal(t, "Fabio", user.Name)
	assert.Nil(t, user.Validate())
	assert.Nil(t, errPassword)
}

func TestUser_WithoudAvatar(t *testing.T) {
	user := NewUser()
	user.Name = "Fabio"
	user.Email = "fabio@fabio.com"
	errPassword := user.SetPassword("fabio10")
	user.DefineAvatar("")

	assert.Equal(t, "", user.AvatarURL)
	assert.NotNil(t, user)
	assert.Nil(t, errPassword)
}

func TestUser_Invalid(t *testing.T) {
	user := NewUser()
	user.Name = "Fabio"
	errPassword := user.SetPassword("")
	user.DefineAvatar("image1.jpg")

	assert.EqualError(t, user.Validate(), "Field 'Email' invalid: required")
	assert.NotNil(t, user.Validate())
	assert.Nil(t, errPassword)
}

func TestUser_InvalidPassword(t *testing.T) {
	user := NewUser()
	user.Name = "Fabio"
	user.Email = "fabio@fabio.com"
	errPassword := user.SetPassword("kSzLdTJAZycMCDfjbWve2RRWaBKnsayujmxCrPXMzTjkfEb7FquhMBfmvU9fXaXPNXbMeUVmkSDfwdZXb3YP9XXshTURUj9v7yj7")
	user.DefineAvatar("image1.jpg")

	assert.NotNil(t, user.Validate())
	assert.NotNil(t, errPassword)
	assert.Equal(t, errPassword.(*errorformatted.ErrorFormatted).Code, errorcode.ErrorUserBcryptError)
}
