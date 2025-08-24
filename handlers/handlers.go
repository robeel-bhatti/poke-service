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

// GetPokemon gets a collection of pokemon, applying the provided query params.
func (ph *PokeHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	reqId := r.Header.Get(constants.RequestIdKey)
	ph.logger.Info("Request received to get a collection of pokemon: ", "params", qp)

	res, err := ph.pokeService.GetPokemon(qp)

	if err != nil {
		ph.logger.Error(err.Error())
		appErr := errors.CreateErrorResponse(r.URL.Path, reqId, err)
		JsonEncode(w, appErr.Status, reqId, appErr)
	} else {
		JsonEncode(w, http.StatusOK, reqId, res)
	}
}

// GetPokemonByName get a pokemon by the name provided in the path variable
func (ph *PokeHandler) GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	n := r.PathValue("name")
	ph.logger.Info("Request received to get the following Pokemon:", "name", n)
	res, err := ph.pokeService.GetPokemonByName(n)
	id := r.Header.Get(constants.RequestIdKey)

	if err != nil {
		ph.logger.Error(err.Error())
		appErr := errors.CreateErrorResponse(r.URL.Path, id, err)
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
}
