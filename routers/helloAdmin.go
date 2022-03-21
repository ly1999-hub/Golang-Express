package routers

import (
	//"fmt"
	"fmt"
	"myapp/models"
	"net/http"

	//"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func HelloAdmin(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username := claims["username"].(string)
	admin := claims["admin"].(bool)
	message := fmt.Sprintf("hello %s is admin %v", username, admin)

	x := &models.X{
		Text: message,
	}

	return c.JSON(http.StatusOK, x)
}
