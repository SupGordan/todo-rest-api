package api

import (
	"net/http"
	"todo-rest-api/pkg/app"
	"todo-rest-api/pkg/e"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	app := app.Gin{C: c}

	app.Response(http.StatusOK, e.SUCCESS, map[string]bool{
		"pong": true,
	})
}
