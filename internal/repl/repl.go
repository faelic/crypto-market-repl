package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type REPL struct{}

var supportedCoins = map[string]bool{
	"bitcoin":  true,
	"ethereum": true,
	"solana":   true,
}

func (r REPL) Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		if input == "/exit" {
			fmt.Println("see you later")
			return
		}

		r.handleInput(input)

	}
}

func (r REPL) handleInput(input string) {
	command, args := parseInput(input)
	_ = args

	switch command {
	case "/help":
		fmt.Println(`Available commands:
	/help
	/list
	/price <coin>
	/market <coin>
	/exit

	Supported coins:
	bitcoin
	ethereum
	solana`)

		return
	case "/price":
		r.handlePrice(args)
		return
	default:
		fmt.Printf("unknown command: %s\n", command)

	}
}

func (r REPL) handlePrice(args []string) {
	message := `usage: /price <coin>
examples: /price bitcoin, /price ethereum, /price solana
`
	if len(args) != 1 {
		fmt.Println(message)
		return
	}

	coin := strings.ToLower(args[0])
	_, ok := supportedCoins[coin]

	if !ok {
		fmt.Printf("unsupported coin: %s\n", coin)
		fmt.Println("supported coins: bitcoin, ethereum, solana")
		return
	}

	fmt.Printf("fetching price for %s ...\n", coin)
}
