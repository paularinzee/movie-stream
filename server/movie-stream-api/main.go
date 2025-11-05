package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paularinzee/server/movie-stream-api/database"
	"github.com/paularinzee/server/movie-stream-api/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️  Warning: .env file not found, using environment variables")
	}

	// Initialize router
	router := gin.Default()

	// Simple test route
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	// --- Configure CORS ---
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	var origins []string

	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
			log.Println("✅ Allowed Origin:", origins[i])
		}
	} else {
		origins = []string{"http://localhost:5173"}
		log.Println("⚠️ Default Allowed Origin: http://localhost:5173")
	}

	corsConfig := cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))
	router.Use(gin.Logger())

	// --- Connect to MongoDB ---
	client := database.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ Failed to reach MongoDB: %v", err)
	}

	log.Println("✅ Successfully connected to MongoDB")

	// --- Register Routes ---
	routes.UserRoutes(router, client)
	routes.MovieRoutes(router, client)

	// --- Graceful Shutdown ---
	srvErrChan := make(chan error)
	go func() {
		log.Println("🚀 Server running on http://localhost:8080")
		srvErrChan <- router.Run(":8080")
	}()

	// Listen for OS signals to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("🛑 Shutting down server...")

	case err := <-srvErrChan:
		if err != nil {
			log.Fatalf("❌ Server error: %v", err)
		}
	}

	// Disconnect MongoDB
	if err := client.Disconnect(ctx); err != nil {
		log.Printf("❌ Error disconnecting from MongoDB: %v", err)
	} else {
		log.Println("✅ Disconnected from MongoDB")
	}

	log.Println("👋 Server gracefully stopped")
}
