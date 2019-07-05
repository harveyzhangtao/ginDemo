package loadredis

import (
	"fmt"
	"ginDemo/pkg/setting"
	"github.com/go-redis/redis"
	"sync"
)

var Client *redis.Client
var once sync.Once

func init() {
	once.Do(func() {
		Client = redis.NewClient(&redis.Options{
			Addr:     setting.RedisSeeting.RedisAddress,
			Password: setting.RedisSeeting.RedisPassword, // no password set
			DB:       setting.RedisSeeting.RedisDB,       // use default DB
			//PoolSize:5,   //默认10*cpu数
		})
	})
}

func ReturnRedisClient() *redis.Client {
	fmt.Println(Client)
	return Client
}