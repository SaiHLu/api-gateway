package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SaiHLu/api-gateway/config"
	"github.com/SaiHLu/api-gateway/internal/transport/gin/handler"
	"github.com/SaiHLu/api-gateway/internal/transport/gin/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type AppServer struct {
	envConfig  *config.EnvConfig
	ginEngine  *gin.Engine
	httpServer *http.Server

	quit chan os.Signal

	// Dependencies Injection
	swaggerHandler     *handler.SwaggerHandler
	healthCheckHandler *handler.HealthCheckHandler

	productClientConn *grpc.ClientConn
	productHandler    *handler.ProductHandler

	orderClientConn *grpc.ClientConn
	orderHandler    *handler.OrderHandler
}

func NewAppServer(envConfig *config.EnvConfig) *AppServer {
	gin.SetMode(envConfig.Mode)
	engine := gin.Default()

	engine.Use(gin.Recovery(), middleware.CorsMiddleware(envConfig))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", envConfig.Port),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	return &AppServer{
		envConfig:  envConfig,
		quit:       quit,
		ginEngine:  engine,
		httpServer: httpServer,
	}
}

func (s *AppServer) Start() error {
	if err := s.dependenciesInjection(); err != nil {
		return fmt.Errorf("failed to inject dependencies: %w", err)
	}

	// Start the HTTP server
	return s.init()
}

func (s *AppServer) init() error {
	errg, ctx := errgroup.WithContext(context.Background())

	errg.Go(func() error {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("failed to start HTTP server: %w", err)
		}

		return nil
	})

	errg.Go(func() error {
		<-s.quit
		log.Println("Shutting down server...")

		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
		defer shutdownCancel()

		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to close HTTP server: %w", err)
		}

		if err := s.orderClientConn.Close(); err != nil {
			return fmt.Errorf("failed to close gRPC license client connection: %w", err)
		}

		if err := s.productClientConn.Close(); err != nil {
			return fmt.Errorf("failed to close gRPC product client connection: %w", err)
		}

		log.Println("Server gracefully stopped")
		return nil
	})

	return errg.Wait()
}
