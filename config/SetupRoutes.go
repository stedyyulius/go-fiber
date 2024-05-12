package config

import (
	"go-fiber/controllers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(db *mongo.Database) {
    app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected")
	})

    app.Get("/statements", func(c *fiber.Ctx) error {
		return controllers.GetDailyStatements(c, db.Collection("Statement"))
	})

    app.Listen(":3000")
}