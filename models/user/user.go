package user

import (
	"todo-rest-api/models"
	"todo-rest-api/pkg/crypt"
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
