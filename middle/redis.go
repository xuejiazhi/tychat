package middle

import (
	"fmt"
	"github.com/go-redis/redis"
	"tychat/util"
)

func InitRedis() {
	//redis addr
	addr := fmt.Sprintf("%s:%d", util.GlobalCfg.Redis.IP, util.GlobalCfg.Redis.Port)
	//连接redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "openIM123",
		DB:       0})
	//defer 关闭
	defer RedisClient.Close()

	//ping
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		util.HandleError(err)
	}

	//
	fmt.Printf("Redis ping %s success!\n", pong)
	//return
	return
}
