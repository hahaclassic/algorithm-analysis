package storage

import (
	"context"
	"errors"

	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/models"
)

var (
	ErrStorageConnection = errors.New("storage: failed to connect to the database")
	ErrDisconnect        = errors.New("storage: failed to disconnect from database")

	ErrClearCollection = errors.New("storage: failed to clear collection")
	ErrSaveRecipe      = errors.New("storage: failed to save recipe")
)

type Storage interface {
	SaveRecipe(ctx context.Context, recipe *models.Recipe) error
	Clear(ctx context.Context) error
}
