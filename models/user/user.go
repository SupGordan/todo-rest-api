package user

import (
	"errors"
	"todo-rest-api/models"
	"todo-rest-api/pkg/crypt"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

func AddUser(data map[string]interface{}) error {
	passwordHash, err := crypt.HashPassword(data["password"].(string))
	if err != nil {
		return err
	}
	user := User{
		Email:    data["email"].(string),
		Password: passwordHash,
	}
	if err = models.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func FindUser(email, password string) (User, error) {
	var user User
	err := models.Db.Select("*").Where(User{Email: email}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return User{}, err
	}
	if !crypt.CheckPassword(password, user.Password) {
		return User{}, errors.New("Password not match")
	}
	return user, nil
}
