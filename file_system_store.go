package poker

import "io"

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func NewFileSystemPlayerStore(database io.ReadSeeker) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{
		database: database,
	}
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
