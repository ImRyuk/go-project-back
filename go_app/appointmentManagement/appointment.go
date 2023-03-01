package appointmentManagement

import (
    "github.com/gofiber/fiber/v2"
	"go_app/model"
	"go_app/db"
	"go_app/utils"
    "go_app/permissions"
    "go_app/authentication"
    "time"
)
func ComputeEndDatetime(datetimeStart time.Time, serviceUid string) (time.Time, string) {
    var service model.Service
    client, _ := db.DbConnection()
    defer client.Close()

    client.QueryRow(db.REQ_GET_SERVICE_BY_UID, serviceUid).Scan(
            &service.ServiceUid,
            &service.Name,
            &service.Duration,
            &service.Price,
            &service.StoreUid,
    )

    duration := time.Duration(service.Duration) * time.Minute
    newTime := datetimeStart.Add(duration)
    return newTime, service.StoreUid
}

func CheckAppointmentExist(
    datetimeStart time.Time, datetimeEnd time.Time, storeUid string) bool {
    var appointment model.Appointment
    client, _ := db.DbConnection()
    defer client.Close()
    query, _ := client.Prepare(db.REQ_CHECK_BOOKING_EXISTS)
    err := query.QueryRow(datetimeStart, datetimeStart, storeUid, datetimeStart, datetimeEnd, storeUid).Scan(
        &appointment.AppointmentUid,
    )
    if err == nil {
        // if err != nil: The booking already exists
        return false
    }
    return true
}


func CreateAppointment(c *fiber.Ctx) error {
    var appointment model.Appointment

	client, _ := db.DbConnection()
	defer client.Close()

    authHeader := c.Get("Authorization")

    if authHeader == "" {
            c.Status(401)
            return c.SendString("Authorization header is missing")
    }

    userUidJwt, err := authentication.VerifyJwt(authHeader)
    if err != nil {
        c.Status(401)
        return err
    }

    if err := c.BodyParser(&appointment); err != nil {
        return err
    }
    matchUser := permissions.CheckUserUid(userUidJwt, appointment.UserUid)
    if matchUser == false {
        return c.Status(401).SendString("You don't have this auhtorisation")
    }
    appointmentUid:= utils.GenerateUUID()

    datetimeEnd, storeUid := ComputeEndDatetime(appointment.DatetimeStart, appointment.ServiceUid)
    appointmentExist := CheckAppointmentExist(appointment.DatetimeStart, datetimeEnd, storeUid)
    if appointmentExist == false {
        return c.Status(400).SendString("The booking for this store is not free")
    }

    query, _ := client.Prepare(db.REQ_CREATE_APPOINTMENT)
    _, es := query.Exec(appointmentUid, appointment.DatetimeStart, datetimeEnd,
        appointment.UserUid, appointment.ServiceUid)

    if es != nil {
        c.Status(400).SendString("Error while creating appointment")
    }
    return c.Status(201).JSON(fiber.Map {
            "appointmentUid": appointmentUid,
            "serviceUid": appointment.ServiceUid,
            "datetimeStart": appointment.DatetimeStart,
            "datetimeEnd": datetimeEnd,
            "userUid": appointment.UserUid,
    })
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

    userUidJwt, err := authentication.VerifyJwt(authHeader)
    if err != nil {
        c.Status(401)
        return err
    }
    userUid := c.Params("userUid")
    matchUser := permissions.CheckUserUid(userUidJwt, userUid)
    if matchUser == false {
        return c.Status(401).SendString("You don't have this auhtorisation")
    }
    rows, err := client.Query(db.REQ_GET_APPOINTMENTS_USER, userUid)
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
            &appointment.StoreType,

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

    userUidJwt, err := authentication.VerifyJwt(authHeader)
    if err != nil {
        c.Status(401)
        return err
    }

    storeUid := c.Params("storeUid")

    checkRoleUser := permissions.CheckRoleUserStore(userUidJwt, storeUid)
    if checkRoleUser != true {
        return c.Status(401).SendString("You don't have authorisation, on this store")
    }

    rows, err := client.Query(db.REQ_GET_APPOINTMENTS_STORE, storeUid)

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