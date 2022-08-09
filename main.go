package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-react-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app := fiber.New()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb"))
	if err != nil {
		panic(err)
	}

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(ctx *fiber.Ctx) error {
		var users []models.User
		coll := client.Database("gomongodb").Collection("users")
		results, err := coll.Find(context.TODO(), bson.M{})
		if err != nil {
			return err
		}

		for results.Next(context.TODO()) {
			var user models.User
			err := results.Decode(&user)
			if err != nil {
				return err
			}
			users = append(users, user)
		}

		return ctx.JSON(&fiber.Map{
			"data": users,
		})
	})

	app.Post("/users", func(ctx *fiber.Ctx) error {

		var user models.User
		err := ctx.BodyParser(&user)
		if err != nil {
			return err
		}

		coll := client.Database("gomongodb").Collection("users")
		one, err := coll.InsertOne(context.TODO(),
			bson.D{
				{"name", user.Name},
			})
		if err != nil {
			return err
		}

		return ctx.JSON(&fiber.Map{
			"data": one,
		})
	})

	err = app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server on port 3000")
}
