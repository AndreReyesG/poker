package poker_test

import (
	"io"
	"os"
	"testing"

	"github.com/AndreReyesG/poker"
)

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}

func TestFileSystemStore(t *testing.T) {
	t.Run("league from reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Moka", "Wins": 10},
			{"Name": "Milky", "Wins": 33}]`)
		defer cleanDatabase()

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

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Moka", "Wins": 10},
			{"Name": "Milky", "Wins": 33}]`)
		defer cleanDatabase()

		store := poker.NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Milky")
		want := 33
		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Moka", "Wins": 10},
			{"Name": "Milky", "Wins": 33}]`)
		defer cleanDatabase()

		store := poker.NewFileSystemPlayerStore(database)

		store.RecordWin("Moka")

		got := store.GetPlayerScore("Moka")
		want := 11
		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Moka", "Wins": 10},
			{"Name": "Milky", "Wins": 33}]`)
		defer cleanDatabase()

		store := poker.NewFileSystemPlayerStore(database)

		store.RecordWin("Rorro")
		got := store.GetPlayerScore("Rorro")
		want := 1
		poker.AssertScoreEquals(t, got, want)
	})
}
