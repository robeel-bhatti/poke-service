// Package clients interacts with the Pokemon API.
package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"poke-ai-service/handlers/errors"
	"poke-ai-service/models"
)

// PokeClient struct whose instances invoke the Pokemon API.
type PokeClient struct {
	baseUrl string
	logger  *slog.Logger
	client  *http.Client
}

// NewPokeClient serves as the constructor.
func NewPokeClient(l *slog.Logger, bu string, hc *http.Client) *PokeClient {
	return &PokeClient{
		baseUrl: bu,
		logger:  l,
		client:  hc,
	}
}

// GetPokemonByName invokes the Pokemon API to get a Pokemon by the provided name
// and return the Pokemon metadata in a custom Pokemon struct or an error object if an unexpected failure occurs.
func (pc *PokeClient) GetPokemonByName(name string) (*models.PokemonResponse, error) {
	path, err := url.JoinPath(pc.baseUrl, name)
	if err != nil {
		return nil, fmt.Errorf("%w: error creating URL to call pokemon API: %v", errors.ErrInternalServerError, err)
	}

	res, err := pc.client.Get(path)

	if err != nil {
		return nil, fmt.Errorf("%w: too many redirects or HTTP protocol error: %v", errors.ErrInternalServerError, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request failed: %w", errors.ErrNotFound)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("%w: error reading response body: %v", errors.ErrInternalServerError, err)
	}

	pokeRes, err := pc.unmarshalPokemon(body)

	if err != nil {
		return nil, fmt.Errorf("%w: error deserializing response body: %v", errors.ErrInternalServerError, err)
	}

	return pokeRes, nil
}

// unmarshalPokemon takes a list of bytes representing a Pokemon from the API response
// and deserializes the data into a custom struct and returns it or an error if an unexpected failure occurs.
func (pc *PokeClient) unmarshalPokemon(p []byte) (*models.PokemonResponse, error) {
	pokeRes := &models.PokemonResponse{}
	err := json.Unmarshal(p, pokeRes)
	return pokeRes, err
}
