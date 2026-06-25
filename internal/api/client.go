package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/faelic/crypto-market-repl/internal/model"
)

type Client struct {
	baseURL string
}

func NewClient() Client {
	return Client{
		baseURL: "https://api.coingecko.com/api/v3",
	}
}

func (c Client) GetPrice(coin string) (float64, error) {
	var data map[string]map[string]float64
	fullURL := fmt.Sprintf("%s/simple/price?ids=%s&vs_currencies=usd", c.baseURL, coin)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	coinData, ok := data[coin]
	if !ok {
		return 0, fmt.Errorf("%s not found in response", coin)
	}

	price, ok := coinData["usd"]
	if !ok {
		return 0, fmt.Errorf("price of %s not found in response", coin)
	}

	return price, nil

}

func (c Client) ListCoins() []model.Coin {
	return []model.Coin{
		{Name: "Bitcoin", Symbol: "BTC", Price: 100000, Change24h: 4.25},
		{Name: "Ethereum", Symbol: "ETH", Price: 5000, Change24h: -2.50},
		{Name: "Solana", Symbol: "SOL", Price: 0.53, Change24h: 3.33},
	}
}
