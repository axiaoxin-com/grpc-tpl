// 依赖服务初始化

package main

import (
	"fmt"
	"time"

	"github.com/axiaoxin-com/goutils"
	"github.com/axiaoxin-com/logging"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	// AppID 服务app id
	AppID = 1
)

var (
	// DB mysql 库
	DB *gorm.DB
	// RedisClient redis 客户端
	RedisClient *redis.Client
	// RedisLocker redis 分布式锁
	RedisLocker *redislock.Client
)

// InitLogging logging初始化
func InitLogging() {
	level := viper.GetString("logging.level")
	if level != "" {
		logging.SetLevel(level)
	}
}

// InitGrpc 初始化grpc客户端
func InitGrpc() {
}

// InitDB 初始化数据库
func InitDB() {
	if DB == nil {
		db, err := goutils.GormMySQL(fmt.Sprintf("%s.test_db", viper.GetString("env")))
		if err != nil {
			logging.Fatal(nil, "init gorm mysql error: "+err.Error())
		}
		gormLogger := logging.NewGormLogger(zap.InfoLevel, zap.InfoLevel, time.Millisecond*500)
		DB = db.Session(&gorm.Session{Logger: gormLogger})
	}
}

// InitRedis 初始化Redis
func InitRedis() {
	if RedisClient == nil {
		cli, err := goutils.RedisClient(fmt.Sprintf("%s", viper.GetString("env")))
		if err != nil {
			logging.Fatal(nil, "init redis error: "+err.Error())
		}
		RedisClient = cli
		RedisLocker = redislock.New(cli)
	}
}
