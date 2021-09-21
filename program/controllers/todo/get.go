package todo

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Get(context *gin.Context) {
	body, error := ioutil.ReadAll(context.Request.Body)
	if error != nil {
		fmt.Println("error: ", error)
		return
	}
	fmt.Println("body: ", string(body))
}
