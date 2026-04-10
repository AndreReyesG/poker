package poker_test

import (
	"strings"
	"testing"

	"github.com/AndreReyesG/poker"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Moka", "Wins": 10},
			{"Name": "Milky", "Wins": 33}]`)

		store := poker.NewFileSystemPlayerStore(database)

		got := store.GetLeague()
		want := []poker.Player{
			{Name: "Moka", Wins: 10},
			{Name: "Milky", Wins: 33},
		}
		poker.AssertLeague(t, got, want)

		//read again
		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})
}
