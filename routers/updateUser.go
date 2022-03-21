package routers

import (
	"context"
	"log"
	"myapp/models"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(c echo.Context) error {
	database := models.GetDatabase()

	colUsers := database.Collection("users")
	id1 := c.FormValue("id")
	email := c.FormValue("email")
	id, err := primitive.ObjectIDFromHex(id1)
	if err != nil {
		log.Println(err)
		os.Exit(4)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	update := bson.D{
		{"$set", bson.D{{"email", email}}},
	}
	result, err := colUsers.UpdateByID(ctx, id, update)

	if err != nil {
		log.Println("err update by id ", err)
		os.Exit(5)
	} else {
		log.Println("update ", result.MatchedCount)
	}
	return nil
}
