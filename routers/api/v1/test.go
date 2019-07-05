package v1

import (
	"fmt"
	"ginDemo/pkg/logging"
	"github.com/gin-gonic/gin"
)

func TestIndexHandler(c *gin.Context)  {
	val, _ := c.Get("id")
	fmt.Printf("%T", val)
	logging.Log.Error("test")
}
