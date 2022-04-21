package main

import (
	"context"

	"github.com/axiaoxin-com/grpc-tpl/pb"
)

type implement struct {
	pb.UnimplementedGrpcTplServiceServer
}

// Post Post方法实现
func (i *implement) Post(ctx context.Context, req *pb.PostReq) (*pb.PostRsp, error) {
	rsp := pb.PostRsp{}
	return &rsp, nil
}

// Get Get方法实现
func (i *implement) Get(ctx context.Context, req *pb.GetReq) (*pb.GetRsp, error) {
	rsp := pb.GetRsp{}
	return &rsp, nil
}
