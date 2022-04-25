package door

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//TODO
/* Author zsj
 * Desc 开发启动
 * Return
 */
func Develop() {

	router := gin.Default()
	router.GET("/first", first)
	router.Run(":8080")
}

func first(c *gin.Context) {
	fmt.Println("lmao")
}

func init() {
	fmt.Println("lmao")
}
