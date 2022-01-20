package routers

import (
	api "todo-rest-api/routers/Api"
	"todo-rest-api/routers/api/auth"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", api.GetPing)

	r.POST("/register", auth.PostRegister)

	return r
}
