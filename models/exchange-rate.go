package models

type ExchangeRate struct {
	Name     string
	Currency string
	Rate     float32
	Img      string
}

func NewExchangeRate(name string, currency string, rate float32, img string) *ExchangeRate {
	return &ExchangeRate{
		Name:     name,
		Currency: currency,
		Rate:     rate,
		Img:      img,
	}
}
