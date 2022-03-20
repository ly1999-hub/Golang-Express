package routers

import (
	"context"
	"encoding/json"
	"log"
	"myapp/models"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(c echo.Context) error {
	database := models.GetDatabase()
	colUsers := database.Collection("users")
	cond := bson.D{}

	cursor, err := colUsers.Find(context.TODO(), cond) //collection.Find

	if err != nil {
		log.Println("err Find colUsers ", err)
		os.Exit(1)
	}

	var user []models.Users //khai báo interface theo cấu trúc đã được định nghĩa struc User ở models thông qua biến user
	err = cursor.All(context.Background(), &user)

	if err != nil {
		log.Println("err All get users", err)
		os.Exit(2)
	} else {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		enc := json.NewEncoder(c.Response())
		for _, user := range user {
			if err = enc.Encode(user); err != nil {
				return err
			}

			c.Response().Flush()
		}
	}

	return nil
}
