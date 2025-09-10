package handler

import (
	"fmt"
	"math/rand"
	"net/http"

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

func (p *ProductHandler) registerRoutes() {
	router := p.router.Group("products")

	router.POST("/", p.Create)
	router.GET("/", p.GetProductList)
}

// Create godoc
// @Summary      Create a new product
// @Description  Create a new product with the provided details
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200      {object} map[string]any
// @Failure      400      {object} map[string]any
// @Failure      500      {object} map[string]any
// @Router       /products/ [post]
func (p *ProductHandler) Create(c *gin.Context) {
	response, err := p.grpcClient.CreateProduct(c.Request.Context(), &pb.CreateProductRequest{
		Product: &pb.Product{
			Id:          rand.Uint64(),
			Name:        fmt.Sprintf("Product #%+v", rand.Intn(1_000)),
			Price:       rand.Float32(),
			Description: "Good Job",
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Product)
}

// Create godoc
// @Summary      Create product list
// @Description  Create product list with the provided details
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200      {object} map[string]any
// @Failure      400      {object} map[string]any
// @Failure      500      {object} map[string]any
// @Router       /products/ [get]
func (p *ProductHandler) GetProductList(c *gin.Context) {
	response, err := p.grpcClient.GetProductList(c.Request.Context(), &pb.GetProductListRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, response.Products)
}
