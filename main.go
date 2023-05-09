package main

import (
	"financial-api/config"
	"financial-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// @title Financial Request Handler API
// @version 2.0
// @description An API Endpoint for handling organizational financial reimbursement requests
// @contact.name APP Support
// @contact.email app_support@norainc.org
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func main() {
	app := Setup()
	log.Fatal(app.Listen(":" + config.ENV("PORT")))
	log.Println("Starting application @ http://localhost:" + config.ENV("PORT"))
}

func Setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "https://finance-requests.vercel.app",
		AllowHeaders:  "Authorization, Content-Type, X-Requested-With",
		ExposeHeaders: "Access-Control-Allow-Origin",
		AllowMethods:  "GET, POST, PUT, DELETE, HEAD, OPTIONS",
	}))
	routes.Use(app)
	app.Use(limiter.New())
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.GetRespHeader("no-cache", "false") == "true"
		},
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	return app
}
