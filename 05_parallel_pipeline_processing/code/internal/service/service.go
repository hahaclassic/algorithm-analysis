package service

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/config"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/models"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/storage"
)

type (
	taskChIn  <-chan *models.Task
	taskChOut chan<- *models.Task
	stageFunc func(ctx context.Context, in taskChIn, out taskChOut)
)

type PiplineService struct {
	cfg *config.PiplineConfig
	db  *storage.Storage
}

func New(cfg *config.PiplineConfig, db *storage.Storage) *PiplineService {
	return &PiplineService{
		db: db,
	}
}

func (PiplineService) configureTaskGenerator(f func(ctx context.Context, out taskChOut), out taskChOut) func(ctx context.Context) {
	return func(ctx context.Context) {
		f(ctx, out)
	}
}

func (PiplineService) configureFinalFunc(f func(ctx context.Context, in taskChIn), in taskChIn) func(ctx context.Context) {
	return func(ctx context.Context) {
		f(ctx, in)
	}
}

func (PiplineService) configureStageFunc(handler stageFunc, in taskChIn, out taskChOut, workersQuantity int) func(ctx context.Context) {
	return func(ctx context.Context) {
		wg := &sync.WaitGroup{}
		wg.Add(workersQuantity)

		for range workersQuantity {
			go func(ctx context.Context) {
				defer wg.Done()
				handler(ctx, in, out)
			}(ctx)
		}

		wg.Wait()
		close(out)
	}
}

func (p *PiplineService) Start() {
	var (
		readChIn  = make(chan *models.Task)
		parseChIn = make(chan *models.Task)
		saveChIn  = make(chan *models.Task)
		logChIn   = make(chan *models.Task)
	)

	stages := []func(ctx context.Context){
		p.configureTaskGenerator(p.readFileNames, readChIn),
		p.configureStageFunc(p.readFileContent, readChIn, parseChIn, p.cfg.ReadStageWorkers),
		p.configureStageFunc(p.parseRecipe, parseChIn, saveChIn, p.cfg.ParseStageWorkers),
		p.configureStageFunc(p.saveRecipe, saveChIn, logChIn, p.cfg.SaveStageWorkers),
		p.configureFinalFunc(p.logTask, logChIn),
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(stages))
	ctx := context.Background()

	for _, stage := range stages {
		go func(ctx context.Context) {
			defer wg.Done()
			stage(ctx)
		}(ctx)
	}

	wg.Wait()
}

// Task generator
func (p *PiplineService) readFileNames(ctx context.Context, out taskChOut) {
	defer close(out)

	files, err := os.ReadDir(p.cfg.SourceDir)
	if err != nil {
		slog.Error("readFile", "err", fmt.Errorf("failed to read dir %s: %w", p.cfg.SourceDir, err))
		return
	}

	for _, file := range files {
		select {
		case <-ctx.Done():
			slog.Info("readFile", "info", "context done.")
			return
		default:
		}

		if file.IsDir() {
			continue
		}

		task := &models.Task{
			Created: time.Now(),
			Meta:    filepath.Join(p.cfg.SourceDir, file.Name()),
			Stages: []*models.StageInfo{
				{
					Name:   "readFileContent",
					Queued: time.Now(),
				},
			},
		}
		out <- task
	}
}

// Stage Func
func (p *PiplineService) readFileContent(ctx context.Context, in taskChIn, out taskChOut) {
	defer close(out)

	for task := range in {
		select {
		case <-ctx.Done():
			slog.Info("Context done.")
			return
		default:
		}
		task.Stages[len(task.Stages)-1].Started = time.Now()

		filePath, ok := task.Meta.(string)
		if !ok {
			slog.Error("readFileContent", "err", "type assertion error")
		}

		content, err := os.ReadFile(filePath)
		if err != nil {
			slog.Error("readFileContent", "err",
				fmt.Errorf("type assertion error: %s, %w", filePath, err))
			continue
		}

		task.Meta = string(content)
		task.Stages[len(task.Stages)-1].Finished = time.Now()
		task.Stages = append(task.Stages, &models.StageInfo{
			Name:   "parseRecipe",
			Queued: time.Now(),
		})
		out <- task
	}
}

// Stage Func
func (p *PiplineService) parseRecipe(ctx context.Context, in taskChIn, out taskChOut) {}

// Stage Func
func (p *PiplineService) saveRecipe(ctx context.Context, in taskChIn, out taskChOut) {}

// Final Func
func (PiplineService) logTask(ctx context.Context, in taskChIn) {}
