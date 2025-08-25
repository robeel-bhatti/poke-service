// Package clients interacts with the Pokemon API.
package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"poke-ai-service/internal/handlers/errors"
	"poke-ai-service/internal/models"
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

// GetPokemonByName invokes the pokemon API to get a pokemon with the provided name
func (pc *PokeClient) GetPokemonByName(name string) (*models.PokemonResponse, error) {
	path, err := url.JoinPath(pc.baseUrl, name)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: error creating URL to call pokemon API: %v", errors.ErrInternalServerError, err)
	}

	res, err := pc.client.Get(path)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: too many redirects or HTTP protocol error: %v", errors.ErrInternalServerError, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"http request failed: %w", errors.ErrNotFound)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: error reading response body: %v", errors.ErrInternalServerError, err)
	}

	var pr models.PokemonResponse
	err = json.Unmarshal(body, &pr)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: error deserializing response body: %v", errors.ErrInternalServerError, err)
	}
	return &pr, nil
}

// GetPokemon invokes the pokemon API to get a collection of pokemon.
func (pc *PokeClient) GetPokemon(o, l string) (*models.PaginatedResponse, error) {
	u, err := url.Parse(pc.baseUrl)
	if err != nil {
		return nil, fmt.Errorf(
			"%w. Error creating URL: %v", errors.ErrInternalServerError, err)
	}

	u.RawQuery = url.Values{
		"offset": {o},
		"limit":  {l},
	}.Encode()

	res, err := pc.client.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf(
			"%w: too many redirects or HTTP protocol error: %v", errors.ErrInternalServerError, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"%w. http request did not return 200 response: %v", errors.ErrInternalServerError, res)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: error reading response body: %v", errors.ErrInternalServerError, err)
	}

	var pr models.PaginatedResponse
	err = json.Unmarshal(body, &pr)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: error deserializing response body: %v", errors.ErrInternalServerError, err)
	}
	return &pr, nil
}
