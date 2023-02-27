package authentication

import (
	"fmt"
	"time"
    "strings"
    "github.com/gofiber/fiber/v2"
	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("super-secret-auth-key") // var env

func CreateJwt(uidUser string) (string, error) {
	claims := jwt.MapClaims{
	        "iss": "Ajouter le nom de notre app", // var env
	        "exp": time.Now().Add(time.Hour * 24).Unix(),
	        "uidUser": uidUser,
	}

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenStr, err := token.SignedString([]byte(SECRET))
    if err != nil {
 		return "", err
    }
    return tokenStr, nil
}

func VerifyJwt(authHeader string) (string, error) {
    authHeaderParts := strings.Split(authHeader, " ")
    tokenStr := authHeaderParts[1]
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(SECRET), nil
    })
    if err != nil {
        return "", err
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return "", fmt.Errorf("Token is invalid")
    }
    return claims["uidUser"].(string), nil
}

func VerifyToken(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            c.Status(401)
            return c.SendString("Authorization header is missing")
        }
        authHeaderParts := strings.Split(authHeader, " ")
        token := authHeaderParts[1]

        _, err := VerifyJwt(token)
        if err != nil {
            c.Status(401)
            return err
        }

        return c.Next()
    }

