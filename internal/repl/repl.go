package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/faelic/crypto-market-repl/internal/api"
)

type REPL struct {
	client api.Client
}

func NewREPL(client api.Client) REPL {
	return REPL{
		client: client,
	}
}

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
	case "/list":
		r.handleList()
		return
	case "/market":
		r.handleMarket(args)
	default:
		fmt.Printf("unknown command: %s\n", command)
		return
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
	_ = r.client

	price, err := r.client.GetPrice(coin)
	if err != nil {
		fmt.Println("failed to fetch price")
		return
	}

	if price < 1 {
		fmt.Printf("%s price: $%.4f\n", coin, price)
		return
	}
	fmt.Printf("%s price: $%.2f\n", coin, price)

}

func (r REPL) handleList() {
	coins := r.client.ListCoins()

	fmt.Println("supported coins market overview:")

	for _, coin := range coins {
		if coin.Price < 1 {
			fmt.Printf("%s (%s): $%.4f | 24h: %.2f%%\n", coin.Name, coin.Symbol, coin.Price, coin.Change24h)
			continue
		}

		fmt.Printf("%s (%s): $%.2f | 24h: %.2f%%\n", coin.Name, coin.Symbol, coin.Price, coin.Change24h)

	}
}

func (r REPL) handleMarket(args []string) {
	message := `usage: /market <coin>
examples: /market bitcoin, /market ethereum, /market solana
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

	marketData, err := r.client.GetMarket(coin)
	if err != nil {
		fmt.Printf("could not get market data of %s", coin)
		return
	}

	fmt.Printf("Name: %s\n", marketData.Name)
	fmt.Printf("Current Price: $%.2f\n", marketData.CurrentPrice)
	fmt.Printf("MarketCap: $%.2f\n", marketData.MarketCap)
	fmt.Printf("MarketCapRank: %d\n", marketData.MarketCapRank)
	fmt.Printf("24h Change %.2f%%\n", marketData.Change24h)
}
