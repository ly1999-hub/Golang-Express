package main

import (
	"myapp/models"
	"myapp/routers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	models.MongoDb()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", routers.GetUser)
	e.POST("/create_user", routers.CreateUser)
	e.DELETE("/delete_user", routers.DeleteOneWithId)
	e.Logger.Fatal(e.Start(":1323"))
}
