package main
import (
	"log/slog"
	"os"
)
func initLogger(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	// TextHandler is human-readable for tailing; swap to NewJSONHandler for machine parsing
	handler := slog.NewTextHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Capture everything during REPL development
	})

	slog.SetDefault(slog.New(handler))
	return f, nil
}