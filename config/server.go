package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"poke-ai-service/clients"
	"poke-ai-service/handlers"
	"poke-ai-service/services"
	"time"
)

type appDeps struct {
	Logger  *slog.Logger
	BaseUrl string
}

func StartApp() error {
	envErr := godotenv.Load()
	if envErr != nil {
		return envErr
	}

	logger := LoadLogger()
	deps := &appDeps{
		Logger:  logger,
		BaseUrl: os.Getenv("POKE_URL"),
	}
	mux := NewMux(deps)

	cfg := LoadConfig()
	logger.Info("starting server...", "port", cfg.Port)
	serverErr := http.ListenAndServe(cfg.Port, mux)
	if serverErr != nil {
		return serverErr
	}
	return nil
}

func NewMux(deps *appDeps) *http.ServeMux {
	mux := http.NewServeMux()
	handler := BuildDeps(deps)
	mux.HandleFunc("/pokemon", handler.GetPokemon)
	mux.HandleFunc("/pokemon/{name}", handler.GetPokemonByName)
	return mux
}

func BuildDeps(deps *appDeps) *handlers.PokeHandler {
	pokeClient := clients.NewPokeClient(deps.Logger, deps.BaseUrl, BuildHttpClient(float64(10)))
	pokeService := services.NewPokemonService(deps.Logger, pokeClient)
	pokeHandler := handlers.NewHandler(deps.Logger, pokeService)
	return pokeHandler
}

func BuildHttpClient(t float64) *http.Client {
	return &http.Client{
		Timeout: time.Duration(t) * time.Second,
	}
}
