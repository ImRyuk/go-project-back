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
    var createService model.CreateService
	client, _ := db.DbConnection()
	defer client.Close()

    uuidService := utils.GenerateUUID()

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

    if err := c.BodyParser(&createService); err != nil {
        return err
    }
    checkRoleUser := permissions.CheckRoleUserStore(uidUserJwt, createService.UidStore)
    if checkRoleUser != true {
        return c.Status(401).SendString("Vous n'avez les droits sur ce store")
    }
    query, _ := client.Prepare(db.REQ_CREATE_SERVICE)

    _, es := query.Exec(uuidService, createService.ServiceName, createService.UidStore, createService.Duration, createService.Price)
    if es != nil {
        return c.Status(400).SendString("Erreur lors de la création du service")
    }

    return c.Status(201).SendString("Created service")
}