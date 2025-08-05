package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"poke-ai-service/errors"
	"poke-ai-service/services"
)

type PokeHandler struct {
	Logger      *slog.Logger
	PokeService *services.PokemonService
}

// NewHandler Returns an instance of the PokeHandler struct.
func NewHandler(logger *slog.Logger, ps *services.PokemonService) *PokeHandler {
	return &PokeHandler{logger, ps}
}

func (ph *PokeHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {

}

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
