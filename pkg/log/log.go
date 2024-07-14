package log

import "log/slog"

func Start() {
	slog.Info("Starting server on :8080")
}
