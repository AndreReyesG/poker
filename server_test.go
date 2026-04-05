package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndreReyesG/poker"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Moka's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Moka", nil)
		response := httptest.NewRecorder()

		poker.PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
