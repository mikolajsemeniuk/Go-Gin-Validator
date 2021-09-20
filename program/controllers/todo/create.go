package todo

import (
	"go-validator/program/entities/todo"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func Create(context *gin.Context) {
	var todo todo.Todo
	if error := context.ShouldBindJSON(&todo); error != nil {
		context.JSON(http.StatusBadRequest, error.Error())
		return
	}
	if error := validator.Validate(todo); error != nil {
		errs := error.(validator.ErrorMap)
		context.JSON(http.StatusBadRequest, errs)
		return
	}
	context.JSON(http.StatusAccepted, todo)
}
