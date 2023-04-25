package services

import (
	context "context"
	"try-gcp-grpc/services/pb"
)

type HelloServiceImpl struct {
	pb.UnimplementedHelloServiceServer
}

func (h HelloServiceImpl) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: request.Message,
	}, nil
}

func NewHelloService() pb.HelloServiceServer {
	return HelloServiceImpl{}
}
