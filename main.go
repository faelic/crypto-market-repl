package main

import (
	"github.com/faelic/crypto-market-repl/internal/api"
	"github.com/faelic/crypto-market-repl/internal/repl"
)

func main() {
	client := api.NewClient()
	r := repl.NewREPL(client)
	r.Start()
}
