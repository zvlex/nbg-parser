package nbg

import "time"

type Exchange struct {
	Date       time.Time   `json:"date"`
	Currencies []*Currency `json:"currencies"`
}

type Currency struct {
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Quantity int     `json:"quantity"`
	Rate     float64 `json:"rate"`
	Diff     float64 `json:"diff"`
}
