package clients

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"poke-ai-service/mocks"
	"testing"
	"time"
)

const (
	TEST_POKE_NAME = "charizard"
	EXPECTED_PATH  = "/" + TEST_POKE_NAME
	TEST_TIMEOUT   = 10
)

func TestGetPokemonByName_PokemonReturned(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != EXPECTED_PATH {
			t.Errorf("Expected %s, got %s", EXPECTED_PATH, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(getMockResponse(t))
	}))
	defer server.Close()

	pc := newTestPokeClient(server.URL)
	res, err := pc.GetPokemonByName(TEST_POKE_NAME)

	if err != nil {
		t.Fatalf("Expected error to be nil, got %s", err.Error())
	}

	if res != nil && res.Name != TEST_POKE_NAME {
		t.Errorf("Expected name to be %s, got %s", TEST_POKE_NAME, res.Name)
	}
}

func newTestPokeClient(baseUrl string) *PokeClient {
	return NewPokeClient(
		mocks.TestLogger,
		baseUrl,
		&http.Client{Timeout: time.Duration(TEST_TIMEOUT) * time.Second},
	)
}

func getMockResponse(t *testing.T) []byte {
	filePath := fmt.Sprintf("testdata/api_responses/%s.json", TEST_POKE_NAME)
	res, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Expected error to be nil when loading mock API response, got %s", err.Error())
	}
	return res
}
