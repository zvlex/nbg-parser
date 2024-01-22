package nbg

import "time"

var positionMap = map[string]int{
	"USD": 1,
	"EUR": 2,
}

var defaultPosition int = 100

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
	Position int
}

func (c *Currency) SetPosition() {
	c.Position = defaultPosition

	position, ok := positionMap[c.Code]

	if ok {
		c.Position = position
	}
}
