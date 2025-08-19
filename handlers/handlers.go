// Package handlers handles HTTP requests
package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"poke-ai-service/handlers/errors"
	"poke-ai-service/services"
	"poke-ai-service/util/constants"
)

type PokeHandler struct {
	logger      *slog.Logger
	pokeService *services.PokemonService
}

func NewHandler(l *slog.Logger, ps *services.PokemonService) *PokeHandler {
	return &PokeHandler{l, ps}
}

func (ph *PokeHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	ph.logger.Info("Request received to list of Pokemon")
	ph.pokeService.GetPokemon(r.URL.Query())
	w.Header().Set(constants.ContentTypeKey, constants.ContentTypeValue)
	w.WriteHeader(http.StatusOK)
}

// GetPokemonByName processes HTTP requests at the "/pokemon/{name}" endpoint
func (ph *PokeHandler) GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	n := r.PathValue("name")
	ph.logger.Info("Request received to get the following Pokemon:", "name", n)
	res, err := ph.pokeService.GetPokemonByName(n)
	id := r.Header.Get(constants.RequestIdKey)

	if err != nil {
		ph.logger.Error(err.Error())
		appErr := errors.CreateErrorResponse(err, r.URL.Path, id)
		JsonEncode(w, appErr.Status, id, appErr)
	} else {
		JsonEncode(w, http.StatusOK, id, res)
	}
}

// JsonEncode encodes a custom struct into a response payload
func JsonEncode(w http.ResponseWriter, code int, id string, res any) {
	if err := json.NewEncoder(w).Encode(res); err != nil {
		code = http.StatusInternalServerError
		http.Error(w, err.Error(), code)
	}
	w.Header().Set(constants.RequestIdKey, id)
	w.Header().Set(constants.ContentTypeKey, constants.ContentTypeValue)
	w.WriteHeader(code)
}
