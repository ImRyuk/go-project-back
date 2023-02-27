package storeManagement

import (
	"go_app/model"
	"github.com/gofiber/fiber/v2"
	"go_app/db"
    "go_app/utils"
)

func CreateUserStore(uidUser string, uidStore string) bool {
    client, _ := db.DbConnection()
    defer client.Close()

    query, _ := client.Prepare(db.REQ_CREATE_USER_STORE)

    _, es := query.Exec(uidUser, uidStore)
    if es != nil {
        return false
    }
    return true
}

func CreateStore(c *fiber.Ctx) error {
    var store model.CreateStore

    uuidStore := utils.GenerateUUID()

	client, _ := db.DbConnection()
	defer client.Close()

    if err := c.BodyParser(&store); err != nil {
        return err
    }

    query, _ := client.Prepare(db.REQ_CREATE_STORE)
    _, err := query.Exec(uuidStore ,store.StoreName,
                store.PostCode, store.Address,
                store.City, store.StoreType)

    if err != nil {
        return c.Status(400).SendString("Erreur lors de la création du store")
    }

    createdUSerStore := CreateUserStore(store.UidUser, uuidStore)
    if createdUSerStore != true {
        return c.Status(400).SendString("Erreur lors de la création du user_store")
    }

    return c.Status(201).SendString("Created store")
}