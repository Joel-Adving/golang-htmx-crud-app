package router

import (
	"strconv"

	"github.com/Joel-Adving/testing-stuff/db"
	"github.com/Joel-Adving/testing-stuff/models"
	"github.com/gofiber/fiber/v2"
)

func TodosRouter(r fiber.Router) {
	r.Route("/todos", func(t fiber.Router) {
		t.Get("/", getTodos)
		t.Get("/:id", getTodoById)
		t.Post("/", createTodo)
		t.Put("/:id", updateTodo)
		t.Delete("/:id", deleteTodo)
	})
}

func getTodos(c *fiber.Ctx) error {
	todos := []*models.Todo{}
	db.Context.Order("id desc").Find(&todos)
	if c.Query("template") == "true" {
		return c.Render("partials/todos", fiber.Map{"Todos": todos})
	}
	return c.JSON(todos)
}

func getTodoById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	todo := new(models.Todo)
	db.Context.Find(&todo, id)
	if c.Query("template") == "true" {
		return c.Render("partials/todo", todo)
	}
	return c.JSON(todo)
}

func createTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	db.Context.Create(todo)
	if c.Query("template") == "true" {
		return c.Render("partials/todo", todo)
	}
	return c.JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendString("error")
	}
	body := new(models.Todo)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	todo := new(models.Todo)
	db.Context.Find(&todo, id)
	todo.Completed = body.Completed
	db.Context.Save(&todo)
	if c.Query("template") == "true" {
		return c.Render("partials/todo", todo)
	}
	return c.JSON(todo)
}

func deleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	db.Context.Delete(&models.Todo{}, id)
	return c.SendString("deleted todo with id: " + c.Params("id"))
}
