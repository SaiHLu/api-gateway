package handler

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "SaiHLu/proto/protogen/go/order"
	productpb "SaiHLu/proto/protogen/go/product"
)

type OrderHandler struct {
	grpcClient pb.OrderServiceClient
	router     *gin.Engine
}

func NewOrderHandler(grpcClient pb.OrderServiceClient, router *gin.Engine) *OrderHandler {
	order := &OrderHandler{
		grpcClient: grpcClient,
		router:     router,
	}

	order.registerRoutes()

	return order
}

func (o *OrderHandler) registerRoutes() {
	router := o.router.Group("orders")

	router.POST("/", o.Create)
	router.GET("/", o.GetOrderList)
}

// Create godoc
// @Summary      Create a new order
// @Description  Create a new order with the provided details
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Success      200      {object} map[string]any
// @Failure      400      {object} map[string]any
// @Failure      500      {object} map[string]any
// @Router       /orders/ [post]
func (o *OrderHandler) Create(c *gin.Context) {
	response, err := o.grpcClient.CreateOrder(c.Request.Context(), &pb.CreateOrderRequest{
		Order: &pb.Order{
			Id:           rand.Uint64(),
			CustomerName: fmt.Sprintf("Customer #%+v", rand.Intn(1_000)),
			Products: []*productpb.Product{
				{
					Id:          rand.Uint64(),
					Name:        fmt.Sprintf("Product #%+v", rand.Intn(1_000)),
					Price:       rand.Float32(),
					Description: "Good Job",
				},
				{
					Id:          rand.Uint64(),
					Name:        fmt.Sprintf("Product #%+v", rand.Intn(1_000)),
					Price:       rand.Float32(),
					Description: "Good Job",
				},
				{
					Id:          rand.Uint64(),
					Name:        fmt.Sprintf("Product #%+v", rand.Intn(1_000)),
					Price:       rand.Float32(),
					Description: "Good Job",
				},
			},
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Order)
}

// Create godoc
// @Summary      Get order list
// @Description  Get order list with the provided details
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Success      200      {object} map[string]any
// @Failure      400      {object} map[string]any
// @Failure      500      {object} map[string]any
// @Router       /orders/ [get]
func (o *OrderHandler) GetOrderList(c *gin.Context) {
	response, err := o.grpcClient.GetOrderList(c.Request.Context(), &pb.GetOrderListRequest{})
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusCreated, response.Orders)
}
