package app

import (
	"context"
	"log/slog"

	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/config"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/service"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/storage/mongodb"
)

func Run(conf *config.Config) {
	ctx := context.Background()

	mongo, err := mongodb.New(ctx, &conf.MongoDB)
	if err != nil {
		slog.Error("MONGODB", "err", err)
		return
	}
	defer mongo.Close(ctx)

	if conf.ClearCollection {
		err := mongo.Clear(ctx)
		if err != nil {
			slog.Error("MONGODB", "err", err)
		} else {
			slog.Info("All data deleted")
		}
	}

	service.New(&conf.Pipeline, mongo).Start(ctx)
}
