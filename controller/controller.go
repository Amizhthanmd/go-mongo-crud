package controller

import "go.mongodb.org/mongo-driver/mongo"

type Controller struct {
	UserDB *mongo.Database
}

func InitController(UserDB *mongo.Database) *Controller {
	return &Controller{
		UserDB: UserDB,
	}
}
