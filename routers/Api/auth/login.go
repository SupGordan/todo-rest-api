package auth

import (
	"net/http"
	"todo-rest-api/models/user"
	"todo-rest-api/pkg/app"
	"todo-rest-api/pkg/e"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func PostLogin(c *gin.Context) {
	app := app.Gin{C: c}

	var form LoginForm

	if err := c.ShouldBind(&form); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	user, err := user.FindUser(form.Email, form.Password)

	if err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}

	app.Response(http.StatusOK, e.SUCCESS, user)
}
