package storeManagement

import (
	"go_app/model"
	"github.com/gofiber/fiber/v2"
	"go_app/db"
    "go_app/utils"
    "go_app/authentication"
    "go_app/permissions"
)

func CreateUserStore(userUid string, storeUid string) bool {
    client, _ := db.DbConnection()
    defer client.Close()

    query, _ := client.Prepare(db.REQ_CREATE_USER_STORE)

    _, es := query.Exec(userUid, storeUid)
    if es != nil {
        return false
    }
    return true
}

func CreateStore(c *fiber.Ctx) error {
    var store model.Store
    client, _ := db.DbConnection()
    defer client.Close()

    authHeader := c.Get("Authorization")

    if authHeader == "" {
            c.Status(401)
            return c.SendString("Authorization header is missing")
    }
    userUidJwt, es := authentication.VerifyJwt(authHeader)
    if es != nil {
        c.Status(401)
        return es
    }
    storeUid := utils.GenerateUUID()
    _, isAuthorized := permissions.VerifyRole(userUidJwt, "EMPLOYEE_ROLE")
    if isAuthorized != true {
        return c.Status(401).SendString("You don't have the permission for create a store")
    }

    if err := c.BodyParser(&store); err != nil {
        return err
    }

    query, _ := client.Prepare(db.REQ_CREATE_STORE)
    _, err := query.Exec(storeUid, store.Name,
                store.PostCode, store.Address,
                store.City, store.StoreType)

    if err != nil {
        return c.Status(400).SendString("Erreur lors de la création du store")
    }

    createdUSerStore := CreateUserStore(userUidJwt, storeUid)
    if createdUSerStore != true {
        return c.Status(400).SendString("Erreur lors de la création du userStore")
    }

    return c.Status(201).JSON(fiber.Map {
            "storeUid": store.StoreUid,
            "name": store.Name,
            "postCode": store.PostCode,
            "address": store.Address,
            "city": store.City,
            "storeType": store.StoreType,
    })
}




