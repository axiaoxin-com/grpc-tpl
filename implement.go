package main

import (
	"context"

	grpc_tpl_pb "github.com/axiaoxin-com/grpc-tpl/pb"
)

type implement struct {
	grpc_tpl_pb.UnimplementedGrpcTplServiceServer
}

// Post Post方法实现
func (i *implement) Post(ctx context.Context, req *grpc_tpl_pb.PostReq) (*grpc_tpl_pb.PostRsp, error) {
	rsp := grpc_tpl_pb.PostRsp{}
	return &rsp, nil
}

// Get Get方法实现
func (i *implement) Get(ctx context.Context, req *grpc_tpl_pb.GetReq) (*grpc_tpl_pb.GetRsp, error) {
	rsp := grpc_tpl_pb.GetRsp{}
	return &rsp, nil
}
