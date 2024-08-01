package logger

import (
	"log"
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	logFile, err := os.OpenFile("pkg/logger/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
		return nil
	}

	handler := slog.NewJSONHandler(logFile, nil)
	logger := slog.New(handler)

	return logger
}
