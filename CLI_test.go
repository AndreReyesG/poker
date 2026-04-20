package poker_test

import (
	"strings"
	"testing"

	"github.com/AndreReyesG/poker"
)

func TestCLI(t *testing.T) {
	t.Run("record Moka win from user input", func(t *testing.T) {
		in := strings.NewReader("Moka wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Moka")
	})

	t.Run("record Milky win from user input", func(t *testing.T) {
		in := strings.NewReader("Milky wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Milky")
	})
}
