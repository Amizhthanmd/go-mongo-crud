package main

import (
	"context"
	"go-mongo-crud/controller"
	"go-mongo-crud/db"
	"go-mongo-crud/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	DbClient := db.ConnectDB()
	defer func() {
		if err := DbClient.Disconnect(context.TODO()); err != nil {
			log.Fatal("Error disconnecting MongoDB client:", err)
		}
		log.Println("MongoDB connection closed.")
	}()

	GetUserDB := db.GetUserDatabase(DbClient)
	Controller := controller.InitController(GetUserDB)

	router := gin.Default()

	routes.UserRoutes(router, Controller)

	router.Run(":8080")
}
