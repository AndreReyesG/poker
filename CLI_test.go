package poker_test

import (
	"strings"
	"testing"

	"github.com/AndreReyesG/poker"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Moka wins\n")
	playerStore := &poker.StubPlayerStore{}

	cli := poker.NewCLI(playerStore, in)
	cli.PlayPoker()

	if len(playerStore.WinCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	got := playerStore.WinCalls[0]
	want := "Moka"

	if got != want {
		t.Errorf("didn't record correct winner, got %q, want %q", got, want)
	}
}
