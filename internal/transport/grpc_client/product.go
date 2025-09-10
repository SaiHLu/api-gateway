package grpcclient

import (
	pb "SaiHLu/proto/protogen/go/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductClient(address string) (*grpc.ClientConn, pb.ProductServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewProductServiceClient(conn)

	return conn, client, nil
}
