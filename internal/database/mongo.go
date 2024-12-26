package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/MarcosMorelli/medication-api/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MedicationCollection *mongo.Collection

func Init(config *config.Config) error {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", config.MongoUser, config.MongoPassword, config.MongoHost, config.MongoPort, config.MongoDatabase)

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		slog.Error("Failed to connect to MongoDB:", "err", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		slog.Error("MongoDB ping failed:", "err", err)
		return err
	}

	MedicationCollection = client.Database(config.MongoDatabase).Collection(config.MongoCollection)

	slog.Info("Connected to MongoDB!")
	return nil
}
