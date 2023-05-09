package routes

import (
	_ "financial-api/docs"
	"financial-api/handlers"
	"financial-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Use(app *fiber.App) {
	app.Get("/swagger/*", swagger.New(
		swagger.Config{
			Title:                  "Finance Request API Endpoint",
			DocExpansion:           "none",
			TryItOutEnabled:        false,
			SupportedSubmitMethods: []string{""},
		},
	))
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/logout", handlers.Logout)

	me := api.Group("/me", middleware.Protected())
	me.Get("/", handlers.GetMe)
	me.Get("/mileage", handlers.GetMyMileage)
	me.Get("/check", handlers.GetMyCheckRequests)
	me.Get("/petty_cash", handlers.GetMyPettyCash)
	vehicle := me.Group("/vehicle")
	vehicle.Post("/", handlers.AddVehicle)
	vehicle.Put("/", handlers.EditVehicle)
	vehicle.Delete("/", handlers.RemoveVehicle)

	user := api.Group("/user", middleware.Protected(), middleware.AdminRoute)
	user.Get("/", handlers.GetAllUsers)
	user.Post("/detail", handlers.GetOneUser)
	user.Post("/name", handlers.GetUserName)
	user.Delete("/deactivate", handlers.DeactivateUser)
	user.Post("/mileage", handlers.UserMileage)
	user.Post("/check", handlers.UserCheckRequests)
	user.Post("/petty_cash", handlers.UserPettyCash)

	mileage := api.Group("/mileage", middleware.Protected())
	mileage.Post("/", handlers.CreateMileage)
	mileage.Put("/", handlers.EditMileage)
	mileage.Delete("/", handlers.DeleteMileage)
	mileage.Post("/approve", middleware.AdminRoute, handlers.ApproveMileage)
	mileage.Post("/reject", middleware.AdminRoute, handlers.RejectMileage)
	mileage.Post("/variance", handlers.MileageVariance)
	mileage.Post("/monthly", middleware.AdminRoute, handlers.MonthlyMileage)
	mileage.Post("/detail", handlers.MileageDetail)

	check_req := api.Group("/check", middleware.Protected())
	check_req.Post("/", handlers.CreateCheckRequest)
	check_req.Put("/", handlers.EditCheckRequest)
	check_req.Delete("/", handlers.DeleteCheckRequest)
	check_req.Post("/approve", middleware.AdminRoute, handlers.ApproveCheckRequest)
	check_req.Post("/reject", middleware.AdminRoute, handlers.RejectCheckRequest)
	check_req.Post("/monthly", middleware.AdminRoute, handlers.MonthlyCheckRequests)
	check_req.Post("/detail", handlers.CheckRequestDetail)

	petty_cash := api.Group("/petty_cash", middleware.Protected())
	petty_cash.Post("/", handlers.CreatePettyCash)
	petty_cash.Put("/", handlers.EditPettyCash)
	petty_cash.Delete("/", handlers.DeletePettyCash)
	petty_cash.Post("/approve", middleware.AdminRoute, handlers.ApprovePettyCash)
	petty_cash.Post("/reject", middleware.AdminRoute, handlers.RejectPettyCash)
	petty_cash.Post("/monthly", middleware.AdminRoute, handlers.MonthlyPettyCash)
	petty_cash.Post("/detail", handlers.PettyCashDetail)

	grant := api.Group("/grant")
	grant.Get("/", handlers.GetAllGrants)
	grant.Post("/detail", handlers.GetOneGrant)
	grant.Post("/check", middleware.Protected(), middleware.AdminRoute, handlers.GrantCheckRequests)
	grant.Post("/mileage", middleware.Protected(), middleware.AdminRoute, handlers.GrantMileage)
	grant.Post("/petty_cash", middleware.Protected(), middleware.AdminRoute, handlers.GrantPettyCash)

	errors := api.Group("/error")
	errors.Post("/", handlers.LogError)
}
