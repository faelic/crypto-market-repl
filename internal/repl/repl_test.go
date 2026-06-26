package repl

import "testing"

func TestIsSupportedCoin(t *testing.T) {
	tests := []struct {
		coin string
		want bool
	}{
		{coin: "bitcoin", want: true},
		{coin: "ethereum", want: true},
		{coin: "solana", want: true},
		{coin: "dogecoin", want: false},
	}

	for _, test := range tests {
		got := isSupportedCoin(test.coin)

		if got != test.want {
			t.Errorf("isSupportedCoin(%s) = %v but expects it to be = %v", test.coin, got, test.want)
		}
	}
}

func TestPriceFormat(t *testing.T) {
	tests := []struct {
		price float64
		want  string
	}{
		{price: 0.84, want: "$0.8400"},
		{price: 1.2995, want: "$1.30"},
		{price: 3, want: "$3.00"},
	}

	for _, test := range tests {
		got := formatPrice(test.price)

		if got != test.want {
			t.Errorf("formatPrice(%v) = %q but expected %q", test.price, got, test.want)
		}
	}
}

func TestPriceFormatBoundary(t *testing.T) {
	tests := []struct {
		price float64
		want  string
	}{
		{price: 0.9999, want: "$0.9999"},
		{price: 1.0, want: "$1.00"},
	}

	for _, test := range tests {
		got := formatPrice(test.price)

		if got != test.want {
			t.Errorf("formatPrice(%v) = %q but expected %q", test.price, got, test.want)
		}
	}
}
