package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

// code for analysis

func (p *PipelineService) processTasks(ctx context.Context, tasks []*models.Task) { // 1
	for _, task := range tasks { // 2 - Первый цикл: итерация по списку задач
		// Вложенный цикл: обработка элементов Meta
		metaItems, ok := task.Meta.([]string) // 3
		if !ok {                              // 4
			slog.Error("processTasks", "err", "type assertion error") // 5
			continue                                                  // 6
		} // 7

		for _, item := range metaItems { // 8 - Вложенный цикл: итерация по элементам Meta
			content, err := os.ReadFile(filepath.Join(p.Conf.SourceDir, item)) // 9
			if err != nil {                                                    // 10
				slog.Error("processTasks", "err", fmt.Errorf("failed to read file '%s': %w", item, err)) // 11
			{

			// Дополнительная обработка содержимого файла
			task.Stages = append(task.Stages, &models.StageInfo{ // 14
				Name:   "processItem", // 15
				Queued: time.Now(),    // 16
			}) // 17

			slog.Info("processTasks", "info", fmt.Sprintf("Processed item: %s", item)) // 18
		} // 19
	} // 20
} // 21
