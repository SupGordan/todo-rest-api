package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

type Server struct {
	RunMode  string
	HttpPort string
}

var ServerSetting = &Server{}

func Setup() {
	var err error
	var exists bool

	if err = godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	AppSetting.JwtSecret, exists = os.LookupEnv("JWT_TOKEN")

	if !exists {
		log.Print("No jwt_token found")
	}

	ServerSetting.RunMode, _ = os.LookupEnv("SERVER_MODE")
	ServerSetting.HttpPort, _ = os.LookupEnv("HTTP_PORT")

}
