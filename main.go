package main

import (
	"log/slog"
	"os"
	"poke-ai-service/config"
)

func main() {
	if err := config.StartApp(); err != nil {
		slog.Error("server failed to start", "reason", err)
		os.Exit(1)
	}
}
