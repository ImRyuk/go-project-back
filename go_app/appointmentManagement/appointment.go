package appointmentManagement

import (
    "github.com/gofiber/fiber/v2"
	"go_app/model"
	"go_app/db"
	"go_app/utils"
    "go_app/permissions"
    "go_app/authentication"
)

func CreateAppointment(c *fiber.Ctx) error {
    var appointment model.Appointment

	client, _ := db.DbConnection()
	defer client.Close()

    authHeader := c.Get("Authorization")

    if authHeader == "" {
            c.Status(401)
            return c.SendString("Authorization header is missing")
    }

    uidUserJwt, err := authentication.VerifyJwt(authHeader)
    if err != nil {
        c.Status(401)
        return err
    }

    if err := c.BodyParser(&appointment); err != nil {
        return err
    }
    uidMatch := permissions.CheckUserUid(uidUserJwt, appointment.UidUser)
    if uidMatch == false {
        return c.Status(401).SendString("You don't have this auhtorisation")
    }
    uuidAppointment := utils.GenerateUUID()
    query, _ := client.Prepare(db.REQ_CREATE_APPOINTMENT)

    _, es := query.Exec(
    		uuidAppointment, appointment.DatetimeStart,
     		appointment.UidUser, appointment.UidService)
    if es != nil {
        c.Status(400)
        return es
    }
    return c.Status(201).SendString("Appointement created")
}

func GetAppointmentsUser(c *fiber.Ctx) error {
    var appointments []model.AppointmentsUser
    client, _ := db.DbConnection()
    defer client.Close()
    authHeader := c.Get("Authorization")

    if authHeader == "" {
            c.Status(401)
            return c.SendString("Authorization header is missing")
    }

    uidUserJwt, err := authentication.VerifyJwt(authHeader)
    if err != nil {
        c.Status(401)
        return err
    }
    uidUser := c.Params("uidUser")
    uidMatch := permissions.CheckUserUid(uidUserJwt, uidUser)
    if uidMatch == false {
        return c.Status(401).SendString("You don't have this auhtorisation")
    }
    rows, err := client.Query(db.REQ_GET_APPOINTMENTS_USER, uidUser)
    for rows.Next() {
        var appointment model.AppointmentsUser
        err = rows.Scan(
            &appointment.ServiceName,
            &appointment.Duration,
            &appointment.Price,
            &appointment.DatetimeStart,
            &appointment.StoreName,
            &appointment.City,
            &appointment.Address,
            &appointment.PostCode,
            &appointment.TypeStore,

        )

        if err != nil {
            return err
        }
        appointments = append(appointments, appointment)
    }

    return c.JSON(fiber.Map{
        "appointments": appointments,
    })
}

func GetAppointmentsStore(c *fiber.Ctx) error {
    var appointments []model.AppointmentsStore

    client, _ := db.DbConnection()
    defer client.Close()
    authHeader := c.Get("Authorization")

    if authHeader == "" {
            c.Status(401)
            return c.SendString("Authorization header is missing")
    }

    uidUserJwt, err := authentication.VerifyJwt(authHeader)
    if err != nil {
        c.Status(401)
        return err
    }

    uidStore := c.Params("uidStore")

    checkRoleUser := permissions.CheckRoleUserStore(uidUserJwt, uidStore)
    if checkRoleUser != true {
        return c.Status(401).SendString("You don't have authorisation, on this store")
    }

    rows, err := client.Query(db.REQ_GET_APPOINTMENTS_STORE, uidStore)

    for rows.Next() {
        var appointment model.AppointmentsStore
        err = rows.Scan(
            &appointment.ServiceName,
            &appointment.FirstName,
            &appointment.LastName,
            &appointment.Email,
            &appointment.DatetimeStart,
        )
        if err != nil {
            return err
        }
        appointments = append(appointments, appointment)
    }

    return c.JSON(fiber.Map{
        "appointments": appointments,
    })

}