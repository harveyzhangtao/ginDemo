package v1

import (
	"fmt"
	"ginDemo/pkg/logging"
	"ginDemo/pkg/redis"
	"github.com/gin-gonic/gin"
	"time"
)

func TestIndexHandler(c *gin.Context)  {
	val, _ := c.Get("id")
	fmt.Printf("%T", val)
	logging.Log.Error("test")
	err := loadredis.Client.Set("test", "aaab", 300*time.Second).Err()
	aa := loadredis.Client.Get("test").Val()
	fmt.Println(err, aa)
}
