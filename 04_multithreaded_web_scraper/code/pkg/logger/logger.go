package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(logFilePath string) {
	file, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(file, nil))
	slog.SetDefault(logger)
}
