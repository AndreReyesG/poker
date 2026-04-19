package poker

import "io"

type CLI struct {
	store PlayerStore
	in    io.Reader
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		store: store,
		in:    in,
	}
}

func (cli *CLI) PlayPoker() {
	cli.store.RecordWin("Moka")
}
