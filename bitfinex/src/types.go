package src

type Ticker struct {
	Bid             float64
	BidSize         float64
	Ask             float64
	AskSize         float64
	DailyChange     float64
	DialyChangePerc float64
	LastPrice       float64
	Volume          float64
	High            float64
	Low             float64
}

type Channel struct {
	Channel string
	ChanId  int
	Symbol  string
	Pair    string
	Event   string
}