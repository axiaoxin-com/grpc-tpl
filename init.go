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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// AppID 服务app id
	AppID = 1
)

var (
	// DB mysql 库
	DB *gorm.DB
	// DefaultMySQLMaxIdleConns db默认最大空闲连接数
	DefaultMySQLMaxIdleConns = 2
	// DefaultMySQLMaxOpenConns db默认最大连接数
	DefaultMySQLMaxOpenConns = 10
	// DefaultMySQLConnMaxLifetime db连接默认可复用的最大时间
	DefaultMySQLConnMaxLifetime = time.Hour

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
		dsn := viper.GetString(fmt.Sprintf("mysql.%s.test_db.dsn", viper.GetString("env")))
		logging.Debug(nil, "InitDB load dsn:"+dsn)

		gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logging.NewGormLogger(zap.InfoLevel, zap.DebugLevel, viper.GetDuration("logging.access_logger.slow_threshold")*time.Millisecond)})
		if err != nil {
			logging.Fatal(nil, "InitDB gorm open error:"+err.Error())
		}
		db, err := gormdb.DB()
		if err != nil {
			logging.Fatal(nil, "InitDB gormdb get db error:"+err.Error())
		}
		db.SetMaxIdleConns(DefaultMySQLMaxIdleConns)
		db.SetConnMaxLifetime(DefaultMySQLConnMaxLifetime)
		db.SetMaxOpenConns(DefaultMySQLMaxOpenConns)
		if err := db.Ping(); err != nil {
			logging.Fatal(nil, "InitDB gorm db ping error:"+err.Error())
		}

		DB = gormdb
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
