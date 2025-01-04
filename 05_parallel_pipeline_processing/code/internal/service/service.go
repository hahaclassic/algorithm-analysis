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
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/logger"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/models"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/parser"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/storage"
)

const issueID = 9144

const (
	readFileNames   = "readFileNames"
	readFileContent = "readFileContent"
	parseRecipe     = "parseRecipe"
	saveRecive      = "saveRecipe"
	// logTask         = "logTask"
)

type (
	taskChIn  <-chan *models.Task
	taskChOut chan<- *models.Task
	stageFunc func(ctx context.Context, in taskChIn, out taskChOut)
)

type PipelineService struct {
	Conf   *config.PipelineConfig
	stat   *models.AvgStatInfo
	parser *parser.Parser
	db     storage.Storage
	logger logger.TaskLogger
}

func New(conf *config.PipelineConfig, db storage.Storage) *PipelineService {
	var log logger.TaskLogger
	log, err := logger.NewCSVLogger("../log.csv", []string{readFileContent, parseRecipe, saveRecive})
	if err != nil {
		slog.Error("failed to create csvLogger", "err", err)
		log = &logger.NoLogger{}
	}

	return &PipelineService{
		Conf:   conf,
		db:     db,
		parser: &parser.Parser{},
		stat:   &models.AvgStatInfo{},
		logger: log,
	}
}

func (PipelineService) configureTaskGenerator(f func(ctx context.Context, out taskChOut), out taskChOut) func(ctx context.Context) {
	return func(ctx context.Context) {
		f(ctx, out)
	}
}

func (PipelineService) configureFinalFunc(f func(ctx context.Context, in taskChIn), in taskChIn) func(ctx context.Context) {
	return func(ctx context.Context) {
		f(ctx, in)
	}
}

func (PipelineService) configureStageFunc(handler stageFunc, in taskChIn, out taskChOut, workersQuantity int) func(ctx context.Context) {
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

func (p *PipelineService) Start(ctx context.Context) {
	var (
		readChIn  = make(chan *models.Task, p.Conf.ChanLen)
		parseChIn = make(chan *models.Task, p.Conf.ChanLen)
		saveChIn  = make(chan *models.Task, p.Conf.ChanLen)
		logChIn   = make(chan *models.Task, p.Conf.ChanLen)
	)

	stages := []func(ctx context.Context){
		p.configureTaskGenerator(p.readFileNames, readChIn),
		p.configureStageFunc(p.readFileContent, readChIn, parseChIn, p.Conf.StageWorkers),
		p.configureStageFunc(p.parseRecipe, parseChIn, saveChIn, p.Conf.StageWorkers),
		p.configureStageFunc(p.saveRecipe, saveChIn, logChIn, p.Conf.StageWorkers),
		p.configureFinalFunc(p.logTask, logChIn),
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(stages))

	for _, stage := range stages {
		go func(ctx context.Context) {
			defer wg.Done()
			stage(ctx)
		}(ctx)
	}

	wg.Wait()
	p.stat.Print(p.Conf.StageWorkers, p.Conf.ChanLen)
}

// Task generator
func (p *PipelineService) readFileNames(ctx context.Context, out taskChOut) {
	defer close(out)

	files, err := os.ReadDir(p.Conf.SourceDir)
	if err != nil {
		slog.Error(readFileNames, "err", fmt.Errorf("failed to read dir %s: %w", p.Conf.SourceDir, err))
		return
	}

	for idx, file := range files {
		select {
		case <-ctx.Done():
			slog.Info(readFileNames, "info", "context done.")
			return
		default:
		}

		if file.IsDir() {
			continue
		}

		task := &models.Task{
			ID:      idx,
			Created: time.Now(),
			Meta:    file.Name(),
			Stages: []*models.StageInfo{
				{
					Name:   readFileContent,
					Queued: time.Now(),
				},
			},
		}
		out <- task
	}
}

// Stage Func
func (p *PipelineService) readFileContent(ctx context.Context, in taskChIn, out taskChOut) {
	for task := range in {
		select {
		case <-ctx.Done():
			slog.Info(readFileContent, "info", "context done.")
			return
		default:
		}
		task.Stages[len(task.Stages)-1].Started = time.Now()

		fileName, ok := task.Meta.(string)
		if !ok {
			slog.Error(readFileContent, "err", "type assertion error")
			continue
		}

		content, err := os.ReadFile(filepath.Join(p.Conf.SourceDir, fileName))
		if err != nil {
			slog.Error(readFileContent, "err",
				fmt.Errorf("failed to read file '%s': %w", fileName, err))
			continue
		}

		task.Meta = content
		task.Stages[len(task.Stages)-1].Finished = time.Now()
		task.Stages = append(task.Stages, &models.StageInfo{
			Name:   parseRecipe,
			Queued: time.Now(),
		})
		out <- task
	}
}

// Stage Func
func (p *PipelineService) parseRecipe(ctx context.Context, in taskChIn, out taskChOut) {
	for task := range in {
		select {
		case <-ctx.Done():
			slog.Info(parseRecipe, "info", "context done.")
			return
		default:
		}
		task.Stages[len(task.Stages)-1].Started = time.Now()

		fileContent, ok := task.Meta.([]byte)
		if !ok {
			slog.Error(parseRecipe, "err", "type assertion error")
			continue
		}

		recipe, err := p.parser.ParseHTMLToRecipe(fileContent)
		if err != nil {
			slog.Error(parseRecipe, "err", fmt.Errorf("invalid hmtl in task %d: %w", task.ID, err))
			continue
		}

		recipe.ID = task.ID
		recipe.IssueID = issueID

		task.Meta = recipe
		task.Stages[len(task.Stages)-1].Finished = time.Now()
		task.Stages = append(task.Stages, &models.StageInfo{
			Name:   saveRecive,
			Queued: time.Now(),
		})
		out <- task
	}
}

// Stage Func
func (p *PipelineService) saveRecipe(ctx context.Context, in taskChIn, out taskChOut) {
	for task := range in {
		select {
		case <-ctx.Done():
			slog.Info(saveRecive, "info", "context done.")
			return
		default:
		}
		task.Stages[len(task.Stages)-1].Started = time.Now()

		recipe, ok := task.Meta.(*models.Recipe)
		if !ok {
			slog.Error(saveRecive, "err", "type assertion error")
			continue
		}

		err := p.db.SaveRecipe(ctx, recipe)
		if err != nil {
			slog.Error(saveRecive, "err", err)
			continue
		}

		task.Meta = nil
		task.Stages[len(task.Stages)-1].Finished = time.Now()
		out <- task
	}
}

// Final Func
func (p *PipelineService) logTask(ctx context.Context, in taskChIn) {
	for task := range in {
		task.Destructed = time.Now()
		p.stat.AddTask(task) // Collect stats
		p.logger.Info(task)  // Log task
		//slog.Info(fmt.Sprintf("logTask: task with id %d has been received", task.ID))
	}
}
