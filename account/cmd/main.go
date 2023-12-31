package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nidhal-ferjani/memrizr-hexagonal/internal/adapters/handlers"
)

func main() {
	// you could insert your favorite logger here for structured or leveled logging
	log.Println("Starting Server...")

	router := gin.Default()

	handlers.NewHandler(&handlers.Config{
		Routes: router,
	})

	// apiRoutes := router.Group("/api/v1")
	// {
	// 	apiRoutes.GET("/account", func(ctx *gin.Context) {

	// 		ctx.JSON(http.StatusOK, gin.H{

	// 			"message": "Hello World i' m here",
	// 		})
	// 	})

	// }

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port: %v\n", server.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	log.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)

	}

}
