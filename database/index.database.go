package database

import (
	"fmt"
	"log"
	"simple-crud-go/configs/db_config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	var dsn string

	switch db_config.DB_DRIVER {

	case "mysql":
		// Contoh dsn MySQL
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db_config.DB_USER,
			db_config.DB_PASSWORD,
			db_config.DB_HOST,
			db_config.DB_PORT,
			db_config.DB_NAME,
		)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	case "postgres":
		// Contoh dsn PostgreSQL
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			db_config.DB_HOST,
			db_config.DB_USER,
			db_config.DB_PASSWORD,
			db_config.DB_NAME,
			db_config.DB_PORT,
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	default:
		log.Fatalf("Unsupported DB_DRIVER: %s (gunakan 'mysql' atau 'postgres')", db_config.DB_DRIVER)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully!")
}
