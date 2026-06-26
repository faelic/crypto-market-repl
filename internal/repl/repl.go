package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/faelic/crypto-market-repl/internal/api"
	"github.com/faelic/crypto-market-repl/internal/model"
)

type REPL struct {
	client api.Client
}

func NewREPL(client api.Client) REPL {
	return REPL{
		client: client,
	}
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

func isSupportedCoin(coin string) bool {
	for _, supportedCoin := range model.SupportedCoins {
		if coin == supportedCoin {
			return true
		}
	}
	return false
}

func formatPrice(price float64) string {
	if price < 1 {
		return fmt.Sprintf("$%.4f", price)
	}

	return fmt.Sprintf("$%.2f", price)
}

func (r REPL) handleInput(input string) {
	command, args := parseInput(input)

	switch command {
	case "/help":
		//still hardcoded would be changed
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

	if !isSupportedCoin(coin) {
		fmt.Printf("unsupported coin: %s\n", coin)
		fmt.Println("supported coins: bitcoin, ethereum, solana")
		return
	}

	price, err := r.client.GetPrice(coin)
	if err != nil {
		fmt.Println("failed to fetch price")
		return
	}

	fmt.Printf("%s price: %s\n", coin, formatPrice(price))
}

func (r REPL) handleList() {
	coins, err := r.client.ListCoins()
	if err != nil {
		fmt.Println("failed to fetch coin list")
		return
	}

	fmt.Println("market overview:")

	for _, coin := range coins {
		fmt.Printf("%s (%s): %s | 24h: %.2f%%\n", coin.Name, coin.Symbol, formatPrice(coin.Price), coin.Change24h)
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

	if !isSupportedCoin(coin) {
		fmt.Printf("unsupported coin: %s\n", coin)
		fmt.Println("supported coins: bitcoin, ethereum, solana")
		return
	}

	marketData, err := r.client.GetMarket(coin)
	if err != nil {
		fmt.Printf("could not get market data of %s\n", coin)
		return
	}

	fmt.Printf("Name: %s\n", marketData.Name)
	fmt.Printf("Current Price: $%.2f\n", marketData.CurrentPrice)
	fmt.Printf("Market Cap: $%.2f\n", marketData.MarketCap)
	fmt.Printf("Market Cap Rank: %d\n", marketData.MarketCapRank)
	fmt.Printf("24h Change: %.2f%%\n", marketData.Change24h)
}
