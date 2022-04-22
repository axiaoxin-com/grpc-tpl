module github.com/axiaoxin-com/grpc-tpl

go 1.16

replace github.com/axiaoxin-com/grpc-tpl/pb => ./pb

require (
	github.com/axiaoxin-com/goutils v1.0.30
	github.com/axiaoxin-com/logging v1.2.13
	github.com/bsm/redislock v0.7.2
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/spf13/viper v1.11.0
	go.uber.org/zap v1.21.0
	golang.org/x/net v0.0.0-20220418201149-a630d4f3e7a2 // indirect
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gorm.io/gorm v1.23.4
)
