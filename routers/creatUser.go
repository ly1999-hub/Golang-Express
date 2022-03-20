package routers

import (
	"context"
	"log"
	"myapp/models"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c echo.Context) error {
	database := models.GetDatabase()

	colUsers := database.Collection("users")

	name := c.FormValue("name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	age := c.FormValue("age")

	var user models.Users

	err := colUsers.FindOne(context.TODO(), bson.M{
		"email": email, //{"keyMongo":"value find"}
	}).Decode(&user)
	if err != nil {
		dos := []interface{}{
			bson.D{
				{"name", name},
				{"email", email},
				{"phone", phone},
				{"age", age},
			},
		}
		result, err := colUsers.InsertMany(context.TODO(), dos)
		if err != nil {
			log.Println("error insertmany :", err)
			os.Exit(3)
		} else {
			log.Printf("Number of documents inserted: %d\n", len(result.InsertedIDs))
			log.Println(os.Stderr.Name())
		}
		return c.String(http.StatusOK, "Created many person")
	} else {
		return c.String(http.StatusOK, "email usered")
	}
}
