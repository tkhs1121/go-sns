package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/tkhs1121/go-sns/handler"
	"github.com/tkhs1121/go-sns/service"
)

func init() {
	var err error

	jst, err := time.LoadLocation("Asia/Tokyo")

	if err != nil {
		panic(err)
	}

	time.Local = jst

	err = godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		panic("Error getting .env data!")
	}
}

func main() {
	service.Connect()
	service.Migrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register", handler.Register)
	
	app.Get("/profile", handler.GetRandProfile)
	app.Put("/profile", handler.UpdateRecommendation)

	app.Listen(":1129")
}
