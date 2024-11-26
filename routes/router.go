package routes

import (
	"go-mongo-crud/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, Controller *controller.Controller) {
	router.POST("/users", Controller.CreateUser)
	router.GET("/users", Controller.ListUsers)
	router.GET("/users/:email", Controller.GetUser)
	router.PUT("/users/:id", Controller.UpdateUser)
	router.DELETE("/users/:id", Controller.DeleteUser)
}
