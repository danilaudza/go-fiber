package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbIp := os.Getenv("DB_IP")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	var sqerr error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUser, dbPass, dbIp, dbPort, dbName)
	DB, sqerr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if sqerr != nil {
		panic("Cannot Connect database")
	}

	fmt.Println("Connected to database")
}
