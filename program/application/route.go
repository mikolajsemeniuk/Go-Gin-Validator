package application

import "go-validator/program/controllers/todo"

func Route() *Application {
	router.GET("/todo", todo.Get)
	router.GET("/todo/:id", todo.GetById)
	router.POST("/todo", todo.Create)
	return &Application{}
}
