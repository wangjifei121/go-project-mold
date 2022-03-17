package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	ProjectLog *zap.SugaredLogger //项目日志
	ProjectDB *gorm.DB //项目mysql
	ProjectRedis *redis.Client //项目redis
)
