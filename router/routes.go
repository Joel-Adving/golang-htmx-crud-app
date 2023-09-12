package router

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Route("/", ViewsRouter)

	api := app.Group("/api")
	TodosRouter(api)
}
