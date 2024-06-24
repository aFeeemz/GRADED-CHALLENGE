package handlers

import (
	"context"
	"net/http"

	"graded-challange-1-aFeeemz/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	Repo models.TransactionRepository
}

func (h *Handler) CreateTransaction(c echo.Context) error {
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}

	if err := h.Repo.CreateTransaction(context.Background(), &transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create transaction",
			"ERROR":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, transaction)
}

func (h *Handler) GetTransactions(c echo.Context) error {
	transactions, err := h.Repo.GetAllTransactions(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to get transactions",
			"ERROR":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ID"})
	}

	transaction, err := h.Repo.GetTransactionByID(context.Background(), objID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Transaction not found"})
	}
	return c.JSON(http.StatusOK, transaction)
}

func (h *Handler) UpdateTransaction(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ID"})
	}

	var updatedTransaction models.Transaction
	if err := c.Bind(&updatedTransaction); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request payload"})
	}

	if err := h.Repo.UpdateTransaction(context.Background(), objID, &updatedTransaction); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to update transaction"})
	}

	return c.JSON(http.StatusOK, updatedTransaction)
}

func (h *Handler) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ID"})
	}

	if err := h.Repo.DeleteTransaction(context.Background(), objID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to delete transaction"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Transaction deleted successfully"})
}
