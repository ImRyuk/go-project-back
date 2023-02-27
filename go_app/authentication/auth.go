package authentication


import (
	"go_app/model"
	"github.com/gofiber/fiber/v2"
	"go_app/db"
	"go_app/utils"
)

func Login(c *fiber.Ctx) error {
	var user model.User
	client, _ := db.DbConnection()
	defer client.Close()

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	err := client.QueryRow(db.REQ_LOGIN_USER, user.Email, user.Password).Scan(&user.Uid, &user.Email)
	if err != nil {
		return err
	}

	token, err := CreateJwt(user.Uid)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"uid": user.Uid,
		"email": user.Email,
		"token": token,
	})
}

func Register(c *fiber.Ctx) error {
	client, _ := db.DbConnection()
	defer client.Close()
	var user model.User
	uuidUser := utils.GenerateUUID()
	if err := c.BodyParser(&user); err != nil {
        return err
    }
    query, es := client.Prepare(db.REQ_CREATE_USER)
    if es != nil {
        panic(es.Error())
    }
    _, err := query.Exec(uuidUser, user.FirstName, user.LastName, user.Email, user.Password, user.ROLE)
		if err != nil {
		return err
	}
    return c.JSON(fiber.Map {
			"first_name": user.FirstName,
			"last_name": user.LastName,
			"email": user.Email,
			"ROLE": user.ROLE,
	})
}