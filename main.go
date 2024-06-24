package main

import (
	"fmt"
	"graded-challange-1-aFeeemz/config"
	"graded-challange-1-aFeeemz/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize MongoDB connection
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	fmt.Println("DBNAME:", dbName)
	fmt.Println("URI:", mongoURI)
	config.InitMongoDBConnection(mongoURI, dbName)

	// Echo instance
	e := echo.New()

	// Routes
	routes.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
