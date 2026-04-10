package poker

import "io"

type FileSystemPlayerStore struct {
	database io.Reader
}

func NewFileSystemPlayerStore(database io.Reader) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{
		database: database,
	}
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}
