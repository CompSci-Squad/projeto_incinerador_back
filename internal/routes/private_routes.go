package routes

import (
	"goapi/internal/modules/books"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	api := app.Group("")

	books.PrivateRoutes(api)
}
