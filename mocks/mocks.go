package mocks

import (
	"io"
	"log/slog"
)

var (
	TestLogger = slog.New(slog.NewJSONHandler(io.Discard, nil))
)
