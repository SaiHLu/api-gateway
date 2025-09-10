package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerHandler struct {
	router *gin.Engine
}

func NewSwaggerHandler(router *gin.Engine) *SwaggerHandler {
	handler := &SwaggerHandler{
		router: router,
	}

	handler.registerRoutes()

	return handler
}

func (h *SwaggerHandler) registerRoutes() {
	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.PersistAuthorization(true)))
}
