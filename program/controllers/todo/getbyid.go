package todo

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetById(context *gin.Context) {
	fmt.Println("id: ", context.Param("id"))
	fmt.Println("query: ", context.Query("q"))
}
