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
		return nil, fmt.Errorf("could not get pokemon with name: %s. %w", name, err)
	}
	return p, nil
}

func (ps PokemonService) GetPokemon(qp url.Values) ([]*models.PokeBasic, error) {
	o := qp.Get("offset")
	l := qp.Get("limit")
	pr, err := ps.client.GetPokemon(o, l)
	if err != nil {
		return nil, fmt.Errorf("could not get pokemon collection: %w", err)
	}

	resLen := len(pr.Results)
	res := make([]*models.PokeBasic, 0, resLen)
	ch := make(chan *models.PokeBasic, resLen)

	for _, elem := range pr.Results {
		go ps.getBasicPokemon(elem.Name, ch)
	}

	for i := 0; i < resLen; i++ {
		if val := <-ch; val != nil {
			res = append(res, val)
		}
	}

	return res, nil
}

func (ps PokemonService) getBasicPokemon(name string, ch chan<- *models.PokeBasic) {
	s, err := ps.client.GetPokemonByName(name)
	if err != nil {
		errMsg := fmt.Errorf("could not get pokemon with name: %s. %w", name, err)
		ps.logger.Error(errMsg.Error())
		ch <- nil
	}

	t := models.Types{}
	if len(s.Types) != 0 {
		t.Primary = s.Types[0].Type.Name
		if len(s.Types) > 1 {
			t.Secondary = s.Types[1].Type.Name
		}
	}

	pb := &models.PokeBasic{
		Name:   s.Name,
		Number: s.Id,
		Type:   t,
		Sprite: s.Sprites.OtherSprite.Home.FrontDefault,
	}

	ch <- pb
}
