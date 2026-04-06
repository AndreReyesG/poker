package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AndreReyesG/poker"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := poker.NewInMemoryPlayerStore()
	server := poker.NewPlayerServer(store)

	player := "Milky"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	poker.AssertStatus(t, response.Code, http.StatusOK)

	poker.AssertResponseBody(t, response.Body.String(), "3")
}
