package configs

import (
	"simple-crud-go/configs/app_config"
	"simple-crud-go/configs/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}
