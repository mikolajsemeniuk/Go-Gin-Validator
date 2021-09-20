package application

import "go-validator/program/controllers/todo"

func Route() *Application {
	router.POST("/todo", todo.Create)
	return &Application{}
}
