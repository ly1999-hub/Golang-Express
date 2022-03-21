package mdw

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func BasicAuthLogin(username string, password string, c echo.Context) (bool, error) {
	if username == "nguyen huu ly" && password == "123" {
		c.Set("username", username)
		c.Set("admin", true)
		return true, nil
	}

	if username == "le quang hoang" && password == "12356" {
		c.Set("username", username)
		c.Set("admin", false)
		return true, nil
	}

	return false, nil
}

func IsAdminMdw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		admin := claims["admin"].(bool)
		if admin {
			next(c)
		}
		return echo.ErrUnauthorized
	}

}
