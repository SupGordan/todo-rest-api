package user_service

import "todo-rest-api/models/user"

type UserRegister struct {
	Email    string
	Password string
}

func (u *UserRegister) Register() error {
	userData := map[string]interface{}{
		"email":    u.Email,
		"password": u.Password,
	}
	if err := user.AddUser(userData); err != nil {
		return err
	}
	return nil
}
