package main

import (
	"github.com/gofiber/fiber/v2"
	"go_app/authentication"
	"go_app/storeManagement"
	"go_app/appointmentManagement"
	"go_app/userManagement"
	 "github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {
	app := fiber.New()
	app.Post("/appointment", appointmentManagement.CreateAppointment)
	app.Get("/appointments-store/:storeUid", appointmentManagement.GetAppointmentsStore)
    app.Get("/appointments-user/:userUid", appointmentManagement.GetAppointmentsUser)

    app.Post("/service", storeManagement.CreateService)
    app.Get("/service/:storeUid", storeManagement.GetServices)

    app.Post("/store", storeManagement.CreateStore)

    app.Get("/store/:storeUid", storeManagement.GetStore)
    app.Get("/store", storeManagement.GetStores)
    app.Get("/store-user/:userUid", storeManagement.GetStoreUser)

 	app.Post("/login", authentication.Login)
	app.Post("/register",authentication.Register)
	app.Get("/user/:userUid", userManagement.GetProfileUser)
  // Allow CORS
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
    AllowOrigins:     "*",
    AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
    AllowHeaders:     "",
    AllowCredentials: false,
	}))
	app.Listen(":3000")
}
