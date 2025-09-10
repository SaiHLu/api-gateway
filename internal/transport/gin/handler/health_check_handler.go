package handler

import "github.com/gin-gonic/gin"

type HealthCheckHandler struct {
	router *gin.Engine
}

func NewHealthCheckHandler(router *gin.Engine) *HealthCheckHandler {
	h := &HealthCheckHandler{
		router: router,
	}

	h.registerRoutes()

	return h
}

func (h *HealthCheckHandler) registerRoutes() {
	h.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
