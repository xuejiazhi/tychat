package middle

import (
	"github.com/go-redis/redis"
	"net/rpc"
)

var RpcClient *rpc.Client
var RedisClient *redis.Client
