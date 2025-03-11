package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Valid(t *testing.T) {
	user := NewUser()
	user.Name = "Fabio"
	user.Email = "fabio@fabio.com"
	user.Password = "fabio10"
	user.DefineAvatar("image1.jpg")

	assert.Equal(t, "image1.jpg", user.AvatarURL)
	assert.NotNil(t, user)
	assert.Equal(t, "fabio@fabio.com", user.Email)
	assert.Equal(t, "Fabio", user.Name)
	assert.Equal(t, "fabio10", user.Password)
	assert.Nil(t, user.Validate())
}

func TestUser_WithoudAvatar(t *testing.T) {
	user := NewUser()
	user.Name = "Fabio"
	user.Email = "fabio@fabio.com"
	user.Password = "fabio10"
	user.DefineAvatar("")

	assert.Equal(t, "", user.AvatarURL)
	assert.NotNil(t, user)
}

func TestUser_Invalid(t *testing.T) {
	user := NewUser()
	user.Name = "Fabio"
	user.Email = "fabio@fabio.com"
	user.Password = ""
	user.DefineAvatar("image1.jpg")

	assert.NotNil(t, user.Validate())
	assert.EqualError(t, user.Validate(), "Field 'Password' invalid: required")
}
