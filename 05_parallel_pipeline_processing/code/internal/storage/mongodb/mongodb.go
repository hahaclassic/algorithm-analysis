package mongodb

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/config"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/models"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	collectionName = "recipes"
)

type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func New(ctx context.Context, cfg *config.MongoConfig) (_ *MongoDB, err error) {
	uri := fmt.Sprintf("mongodb://%s", net.JoinHostPort(cfg.Host, cfg.Port))
	opt := options.Client().ApplyURI(uri)
	opt.SetAuth(options.Credential{Username: cfg.User, Password: cfg.Password})

	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", storage.ErrStorageConnection, err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err = client.Ping(ctxPing, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("%w: %w", storage.ErrStorageConnection, err)
	}

	return &MongoDB{
		client:     client,
		collection: client.Database(cfg.DB).Collection(collectionName),
	}, nil
}

func (m *MongoDB) SaveRecipe(ctx context.Context, recipe *models.Recipe) error {
	ctxInsert, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := m.collection.InsertOne(ctxInsert, recipe)
	if err != nil {
		return fmt.Errorf("%w: %w", storage.ErrSaveRecipe, err)
	}

	return nil
}

func (m *MongoDB) Clear(ctx context.Context) error {
	ctxInsert, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, err := m.collection.DeleteMany(ctxInsert, bson.M{})
	if err != nil {
		return fmt.Errorf("%w: %w", storage.ErrSaveRecipe, err)
	}

	slog.Info(fmt.Sprintf("Deleted %d documents from the collection\n", result.DeletedCount))

	return nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	if err := m.client.Disconnect(ctx); err != nil {
		return fmt.Errorf("%w: %w", storage.ErrStorageConnection, err)
	}

	return nil
}
