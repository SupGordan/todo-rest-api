package auth

import (
	"net/http"
	"todo-rest-api/pkg/app"
	"todo-rest-api/pkg/e"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Email           string `form:"email" json:"email" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	Password_repeat string `form:"password_repeat" json:"password_repeat" binding:"required"`
}

func PostRegister(c *gin.Context) {
	app := app.Gin{C: c}

	var form RegisterForm

	if err := c.ShouldBind(&form); err != nil {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
		return
	}
	if form.Password != form.Password_repeat {
		app.Response(http.StatusBadRequest, e.INVALID_PARAMS, "password and password_repeat must be equal")
		return
	}

	app.Response(http.StatusOK, e.SUCCESS, form)
}
