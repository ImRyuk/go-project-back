package storeManagement

import (
	"github.com/gofiber/fiber/v2"
    "go_app/model"
	"go_app/db"
	"go_app/permissions"
    "go_app/utils"
    "go_app/authentication"
)


func CreateService(c *fiber.Ctx) error {
    var service model.Service
	client, _ := db.DbConnection()
	defer client.Close()

    serviceUid := utils.GenerateUUID()

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

    if err := c.BodyParser(&service); err != nil {
        return err
    }
    checkRoleUser := permissions.CheckRoleUserStore(userUidJwt, service.StoreUid)
    if checkRoleUser != true {
        return c.Status(401).SendString("You don't have permission to create a service for this store")
    }
    query, _ := client.Prepare(db.REQ_CREATE_SERVICE)

    _, es := query.Exec(serviceUid, service.Name, service.StoreUid, service.Duration, service.Price)
    if es != nil {
        return c.Status(400).SendString("Error while creating service")
    }
    return c.Status(201).JSON(fiber.Map {
            "serviceUid": serviceUid,
            "storeUid": service.StoreUid,
            "name": service.Name,
            "duration": service.Duration,
            "price": service.Price,
    })
}


func GetServices(c *fiber.Ctx) error {
    var services []model.Service
    client, _ := db.DbConnection()
    defer client.Close()


    storeUid := c.Params("storeUid")

    rows, err := client.Query(db.REQ_GET_SERVICE_BY_STORE, storeUid)

    if err != nil {
        return err
    }

    for rows.Next() {
        var service model.Service
        err = rows.Scan(
            &service.ServiceUid,
            &service.Name,
            &service.Duration,
            &service.Price,
            &service.StoreUid,
        )
        if err != nil {
            return err
        }
        services = append(services, service)
    }

    return c.JSON(fiber.Map{
       "services": services,
    })
}
