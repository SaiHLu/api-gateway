package app

import (
	"github.com/SaiHLu/api-gateway/internal/transport/gin/handler"
	grpcclient "github.com/SaiHLu/api-gateway/internal/transport/grpc_client"
)

func (s *AppServer) dependenciesInjection() error {
	s.swaggerHandler = handler.NewSwaggerHandler(s.ginEngine)

	orderConn, orderClient, err := grpcclient.NewOrderClient(s.envConfig.GrpcServer.OrderService)
	if err != nil {
		return err
	}
	s.orderClientConn = orderConn
	s.orderHandler = handler.NewOrderHandler(orderClient, s.ginEngine)

	productConn, productClient, err := grpcclient.NewProductClient(s.envConfig.GrpcServer.ProductService)
	if err != nil {
		return err
	}
	s.productClientConn = productConn
	s.productHandler = handler.NewProductHandler(productClient, s.ginEngine)

	return nil
}
