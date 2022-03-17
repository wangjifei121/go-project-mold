package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"projectDemo/global"
)

func InitRedis() {
	global.ProjectRedis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})

	_, err := global.ProjectRedis.Ping().Result()
	if err != nil {
		global.ProjectLog.Info(">>>redis连接失败")
	}
	global.ProjectLog.Info(">>>redis连接成功")
}
