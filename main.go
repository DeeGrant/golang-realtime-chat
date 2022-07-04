package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Initializing App ...")
	app := fiber.New()
	//routes.RegisterRoutes(app)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, world!")
	})

	err := app.Listen(":8000")
	if err != nil {
		panic("App failed to initialize.")
	}
	fmt.Println("App started on port 8000")
}
