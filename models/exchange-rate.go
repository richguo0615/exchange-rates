package models

type ExchangeRate struct {
	Name     string  `json:"Name"`
	Currency string  `json:"Currency"`
	Rate     float32 `json:"Rate"`
	Img      string  `json:"Img"`
	Rank     int     `json:"Rank"`
}

func NewExchangeRate(name string, currency string, rate float32, img string, rank int) *ExchangeRate {
	return &ExchangeRate{
		Name:     name,
		Currency: currency,
		Rate:     rate,
		Img:      img,
		Rank:     rank,
	}
}
