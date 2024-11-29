package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "jonh@email.com", "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, user.ID)
	assert.NotNil(t, user.Email)
	assert.NotNil(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "jonh@email.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "jonh@email.com", "password")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("password"))
	assert.False(t, user.ValidatePassword("pass"))
	assert.NotEqual(t, "password", user.Password)
}
