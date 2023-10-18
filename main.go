package main

import (
	"github.com/danilaudza/go-fiber/database"
	"github.com/danilaudza/go-fiber/database/migration"
	"github.com/danilaudza/go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":3001")
}
