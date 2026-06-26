package model

type Coin struct {
	Name      string
	Symbol    string
	Price     float64
	Change24h float64
}

type MarketData struct {
	Name          string
	CurrentPrice  float64
	MarketCap     float64
	MarketCapRank int
	Change24h     float64
}

var SupportedCoins = []string{
	"bitcoin",
	"ethereum",
	"solana",
}
