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

	defer func() {
		if err == nil {
			return
		}

		if err = client.Disconnect(ctx); err != nil {
			slog.Error("mongo", "err",
				fmt.Errorf("%w: %w", storage.ErrDisconnect, err))
		}
	}()

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

// func (m *MongoDB) GetAllRecipes(ctx context.Context) ([]*models.Recipe, error) {
// 	cursor, err := m.collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: %w", storage.ErrGetAllRecipes, err)
// 	}
// 	defer cursor.Close(ctx)

// 	var recipesBSON []bson.M
// 	if err = cursor.All(ctx, &recipesBSON); err != nil {
// 		return nil, fmt.Errorf("%w: %w", storage.ErrGetAllRecipes, err)
// 	}

// 	var recipes []*models.Recipe
// 	for _, recipeData := range recipesBSON {
// 		var recipe models.Recipe
// 		if err := bson.Unmarshal(recipeData, &recipe); err != nil {
// 			log.Println("Ошибка при распаковке данных: ", err)
// 			continue
// 		}
// 		recipeModels = append(recipeModels, &recipe)
// 	}

// 	return recipes, nil

// 	file, err := os.Create("recipes.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "  ") // Для форматирования с отступами
// 	if err := encoder.Encode(recipes); err != nil {
// 		log.Fatal(err)
// 	}
// }

func (m *MongoDB) Close(ctx context.Context) error {
	if err := m.client.Disconnect(ctx); err != nil {
		return fmt.Errorf("%w: %w", storage.ErrStorageConnection, err)
	}

	return nil
}
