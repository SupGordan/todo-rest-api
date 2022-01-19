package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-rest-api/pkg/setting"
	"todo-rest-api/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routers := routers.InitRouter()
	endPoint := fmt.Sprintf(":%s", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routers,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
