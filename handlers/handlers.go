// Package handlers handles HTTP requests
package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"poke-ai-service/handlers/errors"
	"poke-ai-service/services"
)

type PokeHandler struct {
	Logger      *slog.Logger
	PokeService *services.PokemonService
}

func NewHandler(logger *slog.Logger, ps *services.PokemonService) *PokeHandler {
	return &PokeHandler{logger, ps}
}

// GetPokemonByName processes HTTP requests at the "/pokemon/{name}" endpoint
func (ph *PokeHandler) GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	ph.Logger.Info("Request received to get the following Pokemon:", "name", name)
	res, err := ph.PokeService.GetPokemonByName(name)
	if err != nil {
		ph.Logger.Error(err.Error())
		errors.CreateErrorResponse(err, w, r)
	} else {
		json.NewEncoder(w).Encode(res)
	}
}
