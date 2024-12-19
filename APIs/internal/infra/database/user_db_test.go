package database

import (
	"github.com/DiogoFiuza/learning-golang/APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestUser_Create(t *testing.T) {
	// This line create a in-memory database for testing, when the test finish the database is destroyed
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	// This line create the table users in the database
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("userTest", "test@email.com", "password")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var u entity.User
	err = db.First(&u, "name = ?", "userTest").Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, u.ID)
	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.NotNil(t, u.Password)
}

func TestUser_FindByEmail_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("userTest", "test@email.com", "password")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail("test@email.com")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
