package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"poke-ai-service/internal/clients"
	"poke-ai-service/internal/handlers"
	"poke-ai-service/internal/services"
	"time"
)

type appDeps struct {
	logger  *slog.Logger
	baseUrl string
}

// StartApp gathers metadata to start the server
func StartApp() error {
	var err error

	err = godotenv.Load()
	if err != nil {
		return err
	}

	logger := LoadLogger()
	deps := &appDeps{
		logger:  logger,
		baseUrl: os.Getenv("POKE_URL"),
	}
	mux := NewMux(deps)
	cfg := LoadConfig()
	logger.Info("starting server...", "port", cfg.port)

	err = http.ListenAndServe(cfg.port, mux)
	if err != nil {
		return err
	}
	return nil
}

// NewMux creates a new Multiplexer to handle HTTP requests
func NewMux(deps *appDeps) *http.ServeMux {
	mux := http.NewServeMux()
	h := BuildDeps(deps)
	mux.HandleFunc("/pokemon/{name}", h.GetPokemonByName)
	mux.HandleFunc("/pokemon", h.GetPokemon)
	return mux
}

// BuildDeps creates deps the handlers need
func BuildDeps(deps *appDeps) *handlers.PokeHandler {
	pc := clients.NewPokeClient(deps.logger, deps.baseUrl, BuildHttpClient(float64(10)))
	ps := services.NewPokemonService(deps.logger, pc)
	return handlers.NewHandler(deps.logger, ps)
}

// BuildHttpClient creates an http client with configurable timeout settings
func BuildHttpClient(t float64) *http.Client {
	return &http.Client{
		Timeout: time.Duration(t) * time.Second,
	}
}
