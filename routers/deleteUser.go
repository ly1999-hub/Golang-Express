package routers

import (
	"context"
	"log"
	"myapp/models"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOneWithId(c echo.Context) error {
	Database := models.GetDatabase()

	colPerson := Database.Collection("person")
	id1 := c.FormValue("id")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id, _ := primitive.ObjectIDFromHex(id1)
	var filter = bson.M{"_id": id}
	_, err := colPerson.DeleteOne(ctx, filter)

	if err != nil {
		log.Println(err)
	}

	return nil
}
