package grpcclient

import (
	pb "SaiHLu/proto/protogen/go/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewOrderClient(address string) (*grpc.ClientConn, pb.OrderServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewOrderServiceClient(conn)

	return conn, client, nil
}
