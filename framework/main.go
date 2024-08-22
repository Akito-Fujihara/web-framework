package main

import (
	"github.com/Akito-Fujihara/framework/controllers"
	"github.com/Akito-Fujihara/framework/framework"
)

func main() {
	engine := framework.NewEngine()
	router := engine.Router

	router.Get("/lists", controllers.ListsController)
	router.Get("/lists/:list_id", controllers.ListItemController)
	router.Get("/users", controllers.UsersController)
	router.Get("/students", controllers.StudentsController)

	engine.Run()
}
