package user

import "todo-rest-api/models"

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(data map[string]interface{}) error {
	user := User{
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}
	if err := models.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
