package poker_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndreReyesG/poker"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Moka's score", func(t *testing.T) {
		request := newGetScoreRequest("Moka")
		response := httptest.NewRecorder()

		poker.PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Milky's score", func(t *testing.T) {
		request := newGetScoreRequest("Milky")
		response := httptest.NewRecorder()

		poker.PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}
