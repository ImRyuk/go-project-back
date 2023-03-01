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
	app.Get("/appointments-store/:uidStore", appointmentManagement.GetAppointmentsStore)
    app.Get("/appointments-user/:uidUser", appointmentManagement.GetAppointmentsUser)
    app.Post("/service", storeManagement.CreateService)
    app.Post("/store", authentication.VerifyToken, storeManagement.CreateStore)
 	app.Post("/login", authentication.Login)
	app.Post("/register",authentication.Register)
	app.Post("/appointment", appointmentManagement.CreateAppointment)
	app.Get("/user/:uidUser", userManagement.GetProfileUser)
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
