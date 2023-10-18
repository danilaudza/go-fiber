package migration

import (
	"fmt"
	"log"

	"github.com/danilaudza/go-fiber/database"
	"github.com/danilaudza/go-fiber/model"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}
