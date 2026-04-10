package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("problem parsing league, %w", err)
	}
	return league, nil
}
