package pb

import (
	"time"

	"github.com/axiaoxin-com/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// GrpcClient grpc-tpl grpc client
	GrpcClient GrpcTplServiceClient
)

// InitGrpcClient init grpc-tpl grpc client
func InitGrpcClient(addr string) {
	if GrpcClient != nil {
		return
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(30*time.Second))
	if err != nil {
		logging.Fatal(nil, err.Error())
	}
	GrpcClient = NewGrpcTplServiceClient(conn)
}
