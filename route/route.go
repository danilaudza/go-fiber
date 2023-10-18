package route

import (
	"github.com/danilaudza/go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Post("/user", handler.UserHandlerCreate)
	r.Delete("/user/:id", handler.UserHandlerDelete)
}
