package routes

import (
	"context"
	"graded-challange-1-aFeeemz/config"
	"graded-challange-1-aFeeemz/handlers"
	"graded-challange-1-aFeeemz/repository"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	collection, err := config.ConnectionDatabase(context.Background())
	if err != nil {
		e.Logger.Fatal("Failed to connect to database")
	}

	repo := repository.NewMongoTransactionRepository(collection)
	h := &handlers.Handler{Repo: repo}

	// Transaction endpoints
	e.POST("/transactions", h.CreateTransaction)
	e.GET("/all-transactions", h.GetTransactions)
	e.GET("/transactions/:id", h.GetTransactionByID)
	e.PUT("/edit-transactions/:id", h.UpdateTransaction)
	e.DELETE("/delete-transactions/:id", h.DeleteTransaction)
}
