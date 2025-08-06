package clients

import (
	"poke-ai-service/mocks"
	"testing"
)

const (
	TEST_BASE_URL  = "http://testapi.co"
	TEST_POKE_NAME = "Charizard"
	TEST_FULL_URL  = TEST_BASE_URL + "/" + TEST_POKE_NAME
)

func TestGetPokemonByName_PokemonReturned(t *testing.T) {
	pc := newTestPokeClient()
}

func newTestPokeClient() *PokeClient {
	return NewPokeClient(mocks.TestLogger, TEST_BASE_URL)
}
