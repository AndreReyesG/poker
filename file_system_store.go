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
