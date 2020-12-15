package main

import (
	"github.com/Fedena22/Holiday_bucket_tool/db"
	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public/index.html")

	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{

			"admin": "123456",
		},
	}))

	// Or extend your config for customization
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{

			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendFile("./public/unauthorized.html")
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

	app.Static("/admin", "./public/admin.html")
	db.Migrate()
	log.Fatal(app.Listen(":3000"))
}
