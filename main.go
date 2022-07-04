package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/pusher/pusher-http-go"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load environment variables.")
	}

	fmt.Println("Initializing App ...")
	app := fiber.New()
	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}

	//routes.RegisterRoutes(app)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, world!")
	})

	app.Post("/api/messages", func(ctx *fiber.Ctx) error {
		var data map[string]string

		err := ctx.BodyParser(&data)
		if err != nil {
			return err
		}

		return pusherClient.Trigger("chat", "message", data)
	})

	fmt.Println("App starting on port 8000")
	err = app.Listen(":8000")
	if err != nil {
		panic("App failed to initialize.")
	}
}
