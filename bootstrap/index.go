package bootstrap

import (
	"log"
	"simple-crud-go/configs"
	"simple-crud-go/configs/app_config"
	"simple-crud-go/database"
	"simple-crud-go/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	configs.InitConfig()
	database.ConnectDatabase()
	app := gin.Default()
	route.InitRoute(app)
	app.Run(app_config.PORT)
}
