package poker

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func NewFileSystemPlayerStore(database io.Reader) *FileSystemPlayerStore {
	return &FileSystemPlayerStore{
		database: database,
	}
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(f.database).Decode(&league)
	return league
}
