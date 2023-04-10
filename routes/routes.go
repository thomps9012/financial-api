package routes

import (
	"financial-api/handlers"
	"financial-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Use(app *fiber.App) {
	api := app.Group("/api", logger.New())

	seeds := api.Group("/seeds")
	seeds.Post("/", handlers.SeedData)
	seeds.Delete("/", handlers.DeleteSeeds)

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
	user.Get("/detail", handlers.GetOneUser)
	user.Delete("/deactivate", handlers.DeactivateUser)
	user.Get("/mileage", handlers.UserMileage)
	user.Get("/check", handlers.UserCheckRequests)
	user.Get("/petty_cash", handlers.UserPettyCash)

	mileage := api.Group("/mileage", middleware.Protected())
	mileage.Post("/", handlers.CreateMileage)
	mileage.Get("/variance", handlers.MileageVariance)
	mileage.Get("/monthly", middleware.AdminRoute, handlers.MonthlyMileage)
	mileage.Get("/detail", handlers.MileageDetail)
	mileage.Delete("/", handlers.DeleteMileage)
	mileage.Put("/", handlers.EditMileage)
	mileage.Post("/approve", middleware.AdminRoute, handlers.ApproveMileage)
	mileage.Post("/reject", middleware.AdminRoute, handlers.RejectMileage)

	// break point
	check_req := api.Group("/check", middleware.Protected())
	check_req.Post("/", handlers.CreateCheckRequest)
	check_req.Get("/monthly", middleware.AdminRoute, handlers.MonthlyCheckRequests)
	check_req.Get("/detail", handlers.CheckRequestDetail)
	check_req.Put("/", handlers.EditCheckRequest)
	check_req.Delete("/", handlers.DeleteCheckRequest)
	check_req.Post("/approve", middleware.AdminRoute, handlers.ApproveCheckRequest)
	check_req.Post("/reject", middleware.AdminRoute, handlers.RejectCheckRequest)

	petty_cash := api.Group("/petty_cash", middleware.Protected())
	petty_cash.Post("/", handlers.CreatePettyCash)
	petty_cash.Get("/monthly", middleware.AdminRoute, handlers.MonthlyPettyCash)
	petty_cash.Get("/detail", handlers.PettyCashDetail)
	petty_cash.Put("/", handlers.EditPettyCash)
	petty_cash.Delete("/", handlers.DeletePettyCash)
	petty_cash.Post("/approve", middleware.AdminRoute, handlers.ApprovePettyCash)
	petty_cash.Post("/reject", middleware.AdminRoute, handlers.RejectPettyCash)

	grant := api.Group("/grant", middleware.Protected())
	grant.Get("/", handlers.GetAllGrants)
	grant.Get("/detail", handlers.GetOneGrant)
	grant.Get("/check", middleware.AdminRoute, handlers.GrantCheckRequests)
	grant.Get("/mileage", middleware.AdminRoute, handlers.GrantMileage)
	grant.Get("/petty_cash", middleware.AdminRoute, handlers.GrantPettyCash)

	errors := api.Group("/error")
	errors.Post("/", handlers.LogError)
}
