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
	mileage.Get("/:month/:year", middleware.AdminRoute, handlers.MonthlyMileage)
	mileage_record := mileage.Group("/:id")
	mileage_record.Get("/", handlers.MileageDetail)
	mileage_record.Put("/", handlers.EditMileage)
	mileage_record.Delete("/", handlers.DeleteMileage)
	mileage_record.Post("/approve", middleware.AdminRoute, handlers.ApproveMileage)
	mileage_record.Post("/reject", middleware.AdminRoute, handlers.RejectMileage)

	check_req := api.Group("/check", middleware.Protected())
	check_req.Post("/", handlers.CreateCheckRequest)
	check_req.Get("/:month/:year", middleware.AdminRoute, handlers.MonthlyCheckRequests)
	check_req_record := check_req.Group("/:id")
	check_req_record.Get("/", handlers.CheckRequestDetail)
	check_req_record.Put("/", handlers.EditCheckRequest)
	check_req_record.Delete("/", handlers.DeleteCheckRequest)
	check_req_record.Post("/approve", middleware.AdminRoute, handlers.ApproveCheckRequest)
	check_req_record.Post("/reject", middleware.AdminRoute, handlers.RejectCheckRequest)

	petty_cash := api.Group("/petty_cash", middleware.Protected())
	petty_cash.Post("/", handlers.CreatePettyCash)
	petty_cash.Get("/:month/:year", middleware.AdminRoute, handlers.MonthlyPettyCash)
	petty_cash_record := petty_cash.Group("/:id")
	petty_cash_record.Get("/", handlers.PettyCashDetail)
	petty_cash_record.Put("/", handlers.EditPettyCash)
	petty_cash_record.Delete("/", handlers.DeletePettyCash)
	petty_cash_record.Post("/approve", middleware.AdminRoute, handlers.ApprovePettyCash)
	petty_cash_record.Post("/reject", middleware.AdminRoute, handlers.RejectPettyCash)

	grant := api.Group("/grant", middleware.Protected())
	grant.Get("/", handlers.GetAllGrants)
	single_grant := grant.Group("/:id")
	single_grant.Get("/", handlers.GetOneGrant)
	single_grant.Get("/check", middleware.AdminRoute, handlers.GrantCheckRequests)
	single_grant.Get("/mileage", middleware.AdminRoute, handlers.GrantMileage)
	single_grant.Get("/petty_cash", middleware.AdminRoute, handlers.GrantPettyCash)

	errors := api.Group("/error")
	errors.Post("/", handlers.LogError)
}
