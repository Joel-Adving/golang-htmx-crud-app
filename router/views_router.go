package router

import "github.com/gofiber/fiber/v2"

func ViewsRouter(r fiber.Router) {

	r.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Home",
		}, "layouts/root")
	})

	r.Get("/boost", func(c *fiber.Ctx) error {
		return c.Render("boost", fiber.Map{
			"Title": "Boost",
		}, "layouts/root")
	})

}
