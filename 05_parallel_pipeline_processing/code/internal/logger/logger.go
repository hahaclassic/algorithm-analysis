package logger

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/models"
)

type TaskLogger interface {
	Info(task *models.Task)
	Close() error
}

type CSVTaskLogger struct {
	file   *os.File
	writer *csv.Writer
}

func NewCSVLogger(filepath string, stages []string) (*CSVTaskLogger, error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	w := csv.NewWriter(file)
	headers := []string{"TaskID", "Created", "Destructed"}
	for i := range stages {
		headers = append(headers, stages[i]+"Queued", stages[i]+"Started", stages[i]+"Finished")
	}
	err = w.Write(headers)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to write header: %w", err)
	}

	return &CSVTaskLogger{file: file, writer: w}, nil
}

func (c *CSVTaskLogger) Info(task *models.Task) {
	logLine := []string{strconv.Itoa(task.ID), fmt.Sprintf("%d", task.Created.UnixMicro()), fmt.Sprintf("%d", task.Destructed.UnixMicro())}

	for _, stage := range task.Stages {
		logLine = append(logLine, fmt.Sprintf("%d", stage.Queued.UnixMicro()),
			fmt.Sprintf("%d", stage.Started.UnixMicro()), fmt.Sprintf("%d", stage.Finished.UnixMicro()))
	}

	if err := c.writer.Write(logLine); err != nil {
		slog.Error(fmt.Sprintf("failed to write into csv: %v", err))
	}
	c.writer.Flush()
}

func (c *CSVTaskLogger) Close() error {
	c.writer.Flush()
	return c.file.Close()
}

type NoLogger struct{}

func (c *NoLogger) Info(task *models.Task) {}

func (c *NoLogger) Close() error {
	return nil
}
