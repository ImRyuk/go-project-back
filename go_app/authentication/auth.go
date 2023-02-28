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
	bodyPassword := user.Password

	err := client.QueryRow(db.REQ_GET_USER_BY_MAIL, user.Email).Scan(&user.UserUid, &user.Email, &user.FirstName, &user.LastName, &user.ROLE, &user.Password)
	if err != nil {
		return err
	}

	match := utils.CheckPasswordHash(bodyPassword, user.Password)

	if(match == true){
		token, err := CreateJwt(user.UserUid)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"uid": user.UserUid,
			"email": user.Email,
			"token": token,
		})
	}
	c.Status(403)
	return c.SendString("Invalid credentials")
}

func Register(c *fiber.Ctx) error {
	client, _ := db.DbConnection()
	defer client.Close()
	var user model.User
	userUid := utils.GenerateUUID()
	if err := c.BodyParser(&user); err != nil {
        return err
    }
    query, es := client.Prepare(db.REQ_CREATE_USER)
    if es != nil {
        panic(es.Error())
    }

		hashedPassword, _ := utils.HashPassword(user.Password) 

    _, err := query.Exec(userUid, user.FirstName, user.LastName, user.Email, hashedPassword, user.ROLE)
		if err != nil {
		return err
	}
    return c.JSON(fiber.Map {
    	    "userUid": user.UserUid,
			"firstName": user.FirstName,
			"lastName": user.LastName,
			"email": user.Email,
			"ROLE": user.ROLE,
	})
}