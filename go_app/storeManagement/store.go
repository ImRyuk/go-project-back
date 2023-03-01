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


func GetStore(c *fiber.Ctx) error {
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

    storeUid := c.Params("storeUid")
    matchUser := permissions.CheckRoleUserStore(userUidJwt, storeUid)
    if matchUser == false {
        return c.Status(401).SendString("You don't have this authorization")
    }

    err := client.QueryRow(db.REQ_GET_STORE, storeUid).Scan(
            &store.StoreUid,
            &store.Name,
            &store.StoreType,
            &store.City,
            &store.PostCode,
            &store.Address,
    )

    if err != nil {
        return err
    }

    return c.JSON(fiber.Map{
        "storeUid": store.StoreUid,
        "name": store.Name,
        "postCode": store.PostCode,
        "address": store.Address,
        "city": store.City,
        "storeType": store.StoreType,
    })
}

func GetStores(c *fiber.Ctx) error {
    var stores []model.Store

    client, err := db.DbConnection()
    if err != nil {
        return err
    }
    defer client.Close()

    filters := map[string]interface{}{
        "type_store": c.Query("storeType"),
        "name": c.Query("name"),
        "city": c.Query("city"),
        "post_code": c.Query("postCode"),
        "address": c.Query("address"),
    }

    query, args := utils.MakeFilters(filters)
    if len(args) > 0 {
        query = db.REQ_GET_STORES_BY_FILTER + query
    } else {
        query = db.REQ_GET_STORES
    }

    rows, err := client.Query(query, args...)
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var store model.Store
        err = rows.Scan(
            &store.StoreUid,
            &store.Name,
            &store.StoreType,
            &store.City,
            &store.PostCode,
            &store.Address,
        )
        if err != nil {
            return err
        }
        stores = append(stores, store)
    }

    return c.JSON(fiber.Map{
        "stores": stores,
    })
}
