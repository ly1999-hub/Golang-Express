package main

import (
	"myapp/mdw"
	"myapp/models"
	"myapp/routers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	models.MongoDb()

	isLogin := middleware.JWT([]byte("mysecretkey"))
	isAdmin := mdw.IsAdminMdw

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.POST("/login", routers.Login, middleware.BasicAuth(mdw.BasicAuthLogin))
	e.GET("/admin", routers.HelloAdmin, isLogin, isAdmin)

	e.GET("/users", routers.GetUser, isLogin)
	e.POST("/create_user", routers.CreateUser)
	e.DELETE("/delete_user", routers.DeleteOneWithId)
	e.PUT("/update_user_by_id", routers.UpdateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
