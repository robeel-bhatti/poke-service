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
	ph.logger.Info("Request received to get a collection of pokemon: ", "params", qp)
	res, err := ph.pokeService.GetPokemon(qp)
	ph.createResponse(w, r, res, err, http.StatusOK)
}

// GetPokemonByName get a pokemon by the name provided in the path variable.
func (ph *PokeHandler) GetPokemonByName(w http.ResponseWriter, r *http.Request) {
	n := r.PathValue("name")
	ph.logger.Info("Request received to get the following Pokemon:", "name", n)
	res, err := ph.pokeService.GetPokemonByName(n)
	ph.createResponse(w, r, res, err, http.StatusOK)
}

// createResponse takes care of settings appropriate HTTP headers and HTTP status code and serializing
// the appropriate response body for the API response.
func (ph *PokeHandler) createResponse(w http.ResponseWriter, r *http.Request, res any, err error, code int) {
	w.Header().Set(constants.RequestIdKey, r.Header.Get(constants.RequestIdKey))
	w.Header().Set(constants.ContentTypeKey, constants.ContentTypeValue)
	w.WriteHeader(code)

	encoder := json.NewEncoder(w)

	if err != nil {
		ph.logger.Error(err.Error())
		appErr := errors.CreateErrorResponse(r.URL.Path, err)
		w.WriteHeader(appErr.Status)
		encoder.Encode(appErr)
		return
	}

	encoder.Encode(res)
	return
}
