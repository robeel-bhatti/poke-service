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
	Logger      *slog.Logger
	PokeService *services.PokemonService
}

func NewHandler(logger *slog.Logger, ps *services.PokemonService) *PokeHandler {
	return &PokeHandler{logger, ps}
}

func (ph *PokeHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	ph.Logger.Info("Request received to list of Pokemon")
	ph.PokeService.GetPokemon(r.URL.Query())
	w.Header().Set(constants.ContentTypeKey, constants.ContentTypeValue)
	w.WriteHeader(http.StatusOK)
}

// GetPokemonByName processes HTTP requests at the "/pokemon/{name}" endpoint
func (ph *PokeHandler) GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	ph.Logger.Info("Request received to get the following Pokemon:", "name", name)
	res, err := ph.PokeService.GetPokemonByName(name)
	reqId := r.Header.Get(constants.RequestIdKey)

	if err != nil {
		ph.Logger.Error(err.Error())
		appErr := errors.CreateErrorResponse(err, r.URL.Path, reqId)
		JsonEncode(w, appErr.Status, reqId, appErr)
	} else {
		JsonEncode(w, http.StatusOK, reqId, res)
	}
}

// JsonEncode encodes a custom struct into a response payload
func JsonEncode(w http.ResponseWriter, code int, reqId string, res any) {
	if err := json.NewEncoder(w).Encode(res); err != nil {
		code = http.StatusInternalServerError
		http.Error(w, err.Error(), code)
	}
	w.Header().Set(constants.RequestIdKey, reqId)
	w.Header().Set(constants.ContentTypeKey, constants.ContentTypeValue)
	w.WriteHeader(code)
}
