package poker_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndreReyesG/poker"
)

func TestGETPlayers(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{
			"Moka":  20,
			"Milky": 10,
		},
		nil,
	}
	server := poker.NewPlayerServer(&store)

	tests := []struct {
		name               string
		player             string
		expectedHTTPStatus int
		expectedScore      string
	}{
		{
			name:               "returns Moka's score",
			player:             "Moka",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "20",
		},
		{
			name:               "returns Milky's score",
			player:             "Milky",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "10",
		},
		{
			name:               "returns 404 on missing players",
			player:             "Apollo",
			expectedHTTPStatus: http.StatusNotFound,
			expectedScore:      "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := newGetScoreRequest(tt.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			poker.AssertStatus(t, response.Code, tt.expectedHTTPStatus)
			poker.AssertResponseBody(t, response.Body.String(), tt.expectedScore)
		})
	}
}
func TestStoreWins(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := poker.NewPlayerServer(&store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Moka"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		poker.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.WinCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
		}

		if store.WinCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	store := poker.StubPlayerStore{}
	server := poker.NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		poker.AssertStatus(t, response.Code, http.StatusOK)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}
