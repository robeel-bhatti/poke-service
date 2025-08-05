package services

import (
	"fmt"
	"log/slog"
	"poke-ai-service/clients"
	"poke-ai-service/models"
)

type PokemonService struct {
	Logger *slog.Logger
	Client *clients.PokeClient
}

func NewPokemonService(logger *slog.Logger, client *clients.PokeClient) *PokemonService {
	return &PokemonService{
		Logger: logger,
		Client: client,
	}
}

func (ps PokemonService) GetPokemonByName(name string) (*models.PokemonResponse, error) {
	pokemon, err := ps.Client.GetPokemonByName(name)
	if err != nil {
		return nil, fmt.Errorf("could not get pokemon with name: %s. %w", name, err)
	}
	return pokemon, nil
}
