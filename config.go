package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var FiberConfig = fiber.Config{
	JSONEncoder: json.Marshal,
	JSONDecoder: json.Unmarshal,
	Views:       html.New("./views", ".html"),
}

var StaticConfig = fiber.Static{
	Compress: true,
}
