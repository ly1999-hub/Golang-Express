package routers

import (
	"log"
	"myapp/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.Get("username").(string)
	log.Print(username)
	admin := c.Get("admin").(bool)
	log.Print(admin)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()

	t, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		log.Printf("singed token err %v\n", err)
		return err
	}
	return c.JSON(http.StatusOK, &models.ResLogin{
		Token: t,
	})

}
