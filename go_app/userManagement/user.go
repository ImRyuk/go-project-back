package userManagement
import (
    "go_app/model"
    "github.com/gofiber/fiber/v2"
    "go_app/db"
    "go_app/authentication"
    "go_app/permissions"
)

func GetProfileUser(c *fiber.Ctx) error {
    var profile model.User
    client, _ := db.DbConnection()
    defer client.Close()

    authHeader := c.Get("Authorization")

    if authHeader == "" {
            c.Status(401)
            return c.SendString("Authorization header is missing")
    }

    uidUserJwt, es := authentication.VerifyJwt(authHeader)
    if es != nil {
        c.Status(401)
        return es
    }


    uidUser := c.Params("uidUser")
    uidMatch := permissions.CheckUserUid(uidUserJwt, uidUser)
    if uidMatch == false {
        return c.Status(401).SendString("You don't have this auhtorisation")
    }

    err := client.QueryRow(db.REQ_GET_PROFILE_USER, uidUser).Scan(
            &profile.Uid,
            &profile.FirstName,
            &profile.LastName,
            &profile.Email,
            &profile.ROLE,
    )

    if err != nil {
        return err
    }

    return c.JSON(fiber.Map{
        "uid": profile.Uid,
        "firstName": profile.FirstName,
        "lastName": profile.LastName,
        "email": profile.Email,
        "role": profile.ROLE,

    })
}

