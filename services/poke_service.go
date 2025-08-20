// Package services performs biz logic functionality
package services

import (
	"fmt"
	"log/slog"
	"net/url"
	"poke-ai-service/clients"
	"poke-ai-service/models"
)

type PokemonService struct {
	logger *slog.Logger
	client *clients.PokeClient
}

func NewPokemonService(l *slog.Logger, pc *clients.PokeClient) *PokemonService {
	return &PokemonService{
		logger: l,
		client: pc,
	}
}

// GetPokemonByName performs the biz logic to complete this operation
func (ps PokemonService) GetPokemonByName(name string) (*models.PokemonResponse, error) {
	p, err := ps.client.GetPokemonByName(name)
	if err != nil {
		return nil, fmt.Errorf("could not get p with name: %s. %w", name, err)
	}
	return p, nil
}

func (ps PokemonService) GetPokemon(qp url.Values) ([]*models.PokemonResponse, error) {
	o := qp.Get("offset")
	l := qp.Get("limit")
	pr, err := ps.client.GetPokemon(o, l)
	if err != nil {
		return nil, fmt.Errorf("could not get pokemon. %w", err)
	}
	var res []*models.PokemonResponse
	for _, elem := range pr.Results {
		p, err := ps.client.GetPokemonByName(elem.Name)
		if err != nil {
			return nil, fmt.Errorf("could not get pokemon. %w", err)
		}
		res = append(res, p)
	}
	return res, nil
}
