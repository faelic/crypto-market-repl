package api

type Client struct{}

func NewClient() Client {
	return Client{}
}

func (c Client) GetPrice(coin string) float64 {
	switch coin {
	case "bitcoin":
		return 100000.00
	case "ethereum":
		return 2000.00
	case "solana":
		return 0.75
	default:
		return 0
	}
}

func (c Client) ListCoins() []string {
	return []string{"bitcoin", "ethereum", "solana"}
}
