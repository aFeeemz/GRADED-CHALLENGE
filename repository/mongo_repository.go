package repository

import (
	"context"

	"graded-challange-1-aFeeemz/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTransactionRepository struct {
	collection *mongo.Collection
}

func NewMongoTransactionRepository(collection *mongo.Collection) *MongoTransactionRepository {
	return &MongoTransactionRepository{collection}
}

func (repo *MongoTransactionRepository) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	_, err := repo.collection.InsertOne(ctx, transaction)
	return err
}

func (repo *MongoTransactionRepository) GetAllTransactions(ctx context.Context) ([]*models.Transaction, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []*models.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repo *MongoTransactionRepository) GetTransactionByID(ctx context.Context, id primitive.ObjectID) (*models.Transaction, error) {
	var transaction models.Transaction
	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (repo *MongoTransactionRepository) UpdateTransaction(ctx context.Context, id primitive.ObjectID, transaction *models.Transaction) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": transaction}
	_, err := repo.collection.UpdateOne(ctx, filter, update)
	return err
}

func (repo *MongoTransactionRepository) DeleteTransaction(ctx context.Context, id primitive.ObjectID) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
