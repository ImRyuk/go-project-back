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

    userUidJwt, es := authentication.VerifyJwt(authHeader)
    if es != nil {
        c.Status(401)
        return es
    }


    userUid := c.Params("userUid")
    matchUser := permissions.CheckUserUid(userUidJwt, userUid)
    if matchUser == false {
        return c.Status(401).SendString("You don't have this auhtorisation")
    }

    err := client.QueryRow(db.REQ_GET_PROFILE_USER, userUid).Scan(
            &profile.UserUid,
            &profile.FirstName,
            &profile.LastName,
            &profile.Email,
            &profile.ROLE,
    )

    if err != nil {
        return err
    }

    return c.JSON(fiber.Map{
        "userUid": profile.UserUid,
        "firstName": profile.FirstName,
        "lastName": profile.LastName,
        "email": profile.Email,
        "ROLE": profile.ROLE,

    })
}

