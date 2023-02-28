package main

import (
	"github.com/gofiber/fiber/v2"
	"go_app/authentication"
	"go_app/storeManagement"
	"go_app/appointmentManagement"
	"go_app/userManagement"
)


func main() {
	app := fiber.New()
	app.Get("/appointments-store/:storeUid", appointmentManagement.GetAppointmentsStore)
    app.Get("/appointments-user/:userUid", appointmentManagement.GetAppointmentsUser)
    app.Post("/service", storeManagement.CreateService)
    app.Post("/store", storeManagement.CreateStore)
 	app.Post("/login", authentication.Login)
	app.Post("/register",authentication.Register)
	app.Post("/appointment", appointmentManagement.CreateAppointment)
	app.Get("/user/:userUid", userManagement.GetProfileUser)
	app.Listen(":3000")
}
