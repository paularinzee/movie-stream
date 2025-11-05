package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	mongoURI     string
	databaseName string
	envLoaded    bool
)

// init loads environment variables once when the package is imported
func init() {
	loadEnv()
}

// loadEnv loads MongoDB configuration from .env or environment variables
func loadEnv() {
	if envLoaded {
		return
	}

	// Load .env only if it exists
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	mongoURI = os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("❌ MONGODB_URI environment variable is not set")
	}

	databaseName = os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		databaseName = "movie_stream_db"
		log.Printf("⚠️ DATABASE_NAME not set, using default: %s", databaseName)
	}

	log.Println("✅ MongoDB environment configuration loaded")
	envLoaded = true
}

// Connect establishes a MongoDB client connection
func Connect() *mongo.Client {
	loadEnv()

	clientOpts := options.Client().
		ApplyURI(mongoURI).
		SetServerSelectionTimeout(5 * time.Second).
		SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Fatalf("❌ Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
	}

	log.Println("✅ Successfully connected to MongoDB")
	return client
}

// OpenCollection returns a MongoDB collection handle
func OpenCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	if client == nil {
		log.Println("⚠️ MongoDB client is nil")
		return nil
	}
	if collectionName == "" {
		log.Println("⚠️ Collection name cannot be empty")
		return nil
	}

	loadEnv()
	log.Printf("📂 Opening collection: %s in database: %s", collectionName, databaseName)
	return client.Database(databaseName).Collection(collectionName)
}

// Disconnect closes the MongoDB client connection safely
func Disconnect(client *mongo.Client) error {
	if client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		return fmt.Errorf("❌ failed to disconnect MongoDB: %w", err)
	}

	log.Println("✅ MongoDB disconnected successfully")
	return nil
}

// HealthCheck verifies MongoDB connectivity
func HealthCheck(client *mongo.Client) error {
	if client == nil {
		return fmt.Errorf("MongoDB client is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("MongoDB health check failed: %w", err)
	}

	return nil
}

// GetDatabase returns the MongoDB database instance
func GetDatabase(client *mongo.Client) *mongo.Database {
	if client == nil {
		log.Println("⚠️ MongoDB client is nil")
		return nil
	}
	loadEnv()
	return client.Database(databaseName)
}

// GetCollectionNames lists all collection names in the database
func GetCollectionNames(client *mongo.Client) ([]string, error) {
	if client == nil {
		return nil, fmt.Errorf("MongoDB client is nil")
	}

	loadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	names, err := client.Database(databaseName).ListCollectionNames(ctx, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to list MongoDB collections: %w", err)
	}

	return names, nil
}
