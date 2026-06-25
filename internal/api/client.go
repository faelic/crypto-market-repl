package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func (c Client) ListCoins() ([]model.Coin, error) {
	var response []struct {
		Name                    string  `json:"name"`
		Symbol                  string  `json:"symbol"`
		CurrentPrice            float64 `json:"current_price"`
		PriceChangePercentage24 float64 `json:"price_change_percentage_24h"`
	}
	fullURL := fmt.Sprintf("%s/coins/markets?vs_currency=usd&ids=bitcoin,ethereum,solana", c.baseURL)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return []model.Coin{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []model.Coin{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return []model.Coin{}, err
	}

	if len(response) == 0 {
		return []model.Coin{}, fmt.Errorf("could not return any coin from api")
	}

	coins := make([]model.Coin, 0, len(response))

	for _, item := range response {
		coins = append(coins, model.Coin{
			Name:      item.Name,
			Symbol:    strings.ToUpper(item.Symbol),
			Price:     item.CurrentPrice,
			Change24h: item.PriceChangePercentage24,
		})
	}
	return coins, nil
}

func (c Client) GetMarket(coin string) (model.MarketData, error) {
	var response []struct {
		Name                    string  `json:"name"`
		CurrentPrice            float64 `json:"current_price"`
		MarketCap               float64 `json:"market_cap"`
		MarketCapRank           int     `json:"market_cap_rank"`
		PriceChangePercentage24 float64 `json:"price_change_percentage_24h"`
	}

	fullURL := fmt.Sprintf("%s/coins/markets?vs_currency=usd&ids=%s", c.baseURL, coin)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return model.MarketData{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.MarketData{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return model.MarketData{}, err
	}

	if len(response) == 0 {
		return model.MarketData{}, fmt.Errorf("no market data found for %s", coin)
	}

	item := response[0]

	return model.MarketData{
		Name:          item.Name,
		CurrentPrice:  item.CurrentPrice,
		MarketCap:     item.MarketCap,
		MarketCapRank: item.MarketCapRank,
		Change24h:     item.PriceChangePercentage24,
	}, nil
}
