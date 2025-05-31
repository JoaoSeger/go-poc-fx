package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go-poc-fx/internal/user/domain"
	"go-poc-fx/internal/user/presentation"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Server represents the HTTP server
type Server struct {
	router *gin.Engine
	http   *http.Server
}

// New creates a new server instance with routes configured
func New(userService domain.UserService, port string) *Server {
	// Setup Gin router with routes
	router := setupRoutes(userService)

	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		router: router,
		http:   httpServer,
	}
}

// setupRoutes configures all application routes
func setupRoutes(userService domain.UserService) *gin.Engine {
	router := gin.Default()

	// Create user controller
	userController := presentation.NewUserController(userService)

	// Register user routes
	userController.RegisterRoutes(router)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	return router
}

// Start begins listening for HTTP requests
func (s *Server) Start() error {
	fmt.Printf("🚀 Server started on %s\n", s.http.Addr)
	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

// Stop gracefully shuts down the server
func (s *Server) Stop(ctx context.Context) error {
	fmt.Println("🛑 Shutting down server...")
	return s.http.Shutdown(ctx)
}

// StartServer starts the HTTP server as an fx hook
func StartServer(lc fx.Lifecycle, server *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Start(); err != nil {
					fmt.Printf("❌ Error starting server: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop(ctx)
		},
	})
}
