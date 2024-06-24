package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TransactionRepository defines the methods to interact with the database
type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *Transaction) error
	GetAllTransactions(ctx context.Context) ([]*Transaction, error)
	GetTransactionByID(ctx context.Context, id primitive.ObjectID) (*Transaction, error)
	UpdateTransaction(ctx context.Context, id primitive.ObjectID, transaction *Transaction) error
	DeleteTransaction(ctx context.Context, id primitive.ObjectID) error
}
