// Package mocks provides mock deps for app-wide unit tests
package mocks

import (
	"io"
	"log/slog"
)

var (
	TestLogger = slog.New(slog.NewJSONHandler(io.Discard, nil))
)
