package handler

import (
	"github.com/gin-gonic/gin"

	pb "SaiHLu/proto/protogen/go/product"
)

type ProductHandler struct {
	grpcClient pb.ProductServiceClient
	router     *gin.Engine
}

func NewProductHandler(grpcClient pb.ProductServiceClient, router *gin.Engine) *ProductHandler {
	product := &ProductHandler{
		grpcClient: grpcClient,
		router:     router,
	}

	product.registerRoutes()

	return product
}

func (o *ProductHandler) registerRoutes() {}
