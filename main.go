// Package main GRPC服务
package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/axiaoxin-com/goutils"
	grpc_tpl_pb "github.com/axiaoxin-com/grpc-tpl/pb"
	"github.com/axiaoxin-com/logging"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAddr    = flag.String("grpc_addr", ":1888", "The grpc server address")
	gatewayAddr = flag.String("gateway_addr", ":1999", "The grpc gateway address")
	conf        = flag.String("conf", "./config.toml", "The service config file")
)

func init() {
	flag.Parse()

	// init viper
	goutils.InitViper(*conf, nil)
	logging.Info(nil, "init viper config", zap.Any("config", viper.AllSettings()))

	// init logging
	InitLogging()
	// init db
	InitDB()
	// init grpc
	InitGrpc()
	// init redis
	InitRedis()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// gateway
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := grpc_tpl_pb.RegisterGrpcTplServiceHandlerFromEndpoint(ctx, mux, *grpcAddr, opts); err != nil {
		logging.Fatal(ctx, err.Error())
	}
	logging.Infof(ctx, "grpc gateway listening at %v", *gatewayAddr)
	go http.ListenAndServe(*gatewayAddr, mux)

	// grpc
	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logging.Fatal(ctx, err.Error())
	}
	logger := logging.CloneLogger("grpc_tpl_grpcserver").WithOptions(zap.AddCallerSkip(2))
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	grpc_tpl_pb.RegisterGrpcTplServiceServer(server, &implement{})
	logging.Infof(ctx, "grpc server listening at %v", *grpcAddr)
	if err := server.Serve(lis); err != nil {
		logging.Fatal(ctx, err.Error())
	}
}
