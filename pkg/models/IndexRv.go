package models

import (
	"math/rand/v2"
	"strings"
	"time"
)

type InstrumentModel struct {
	IndexID  string `json:"IndexID"`
	MarketID string `json:"MarketID"`
	MICCode  string `json:"MICCode"`
}
type MDFullGrpModel struct {
	PrevPx           float64 `json:"PrevPx"`
	LastTradeDate    string  `json:"LastTradeDate"`
	HighPx52Week     float64 `json:"HighPx52Week"`
	HighPx52WeekDate string  `json:"HighPx52WeekDate"`
	LowPx52Week      float64 `json:"LowPx52Week"`
	LowPx52WeekDate  string  `json:"LowPx52WeekDate"`
}
type IndexRvModel struct {
	TradeDate  string          `json:"TradeDate"`
	Timestamp  time.Time       `json:"Timestamp"`
	Instrument InstrumentModel `json:"Instrument"`
	MDFullGrp  MDFullGrpModel  `json:"MDFullGrp"`
}

func GenerateRandomMessage() IndexRvModel {
	var randomDay = rune(rand.IntN(30))
	seed := rand.NewPCG(0, 3400)
	seed2 := rand.NewPCG(100, 300)
	prevPx := rand.New(seed)
	highPx52Week := prevPx.Float64() + (rand.New(seed2).Float64())
	lowPx52Week := prevPx.Float64() - (rand.New(seed2).Float64())
	return IndexRvModel{
		TradeDate: strings.Join([]string{"2024-11-", string(randomDay)}, ""),
		Timestamp: time.Now(),
		Instrument: InstrumentModel{
			"EQTY",
			"MICCode",
			"XBOG",
		},
		MDFullGrp: MDFullGrpModel{
			prevPx.Float64(),
			"2024-11-08",
			highPx52Week,
			"2024-11-06",
			lowPx52Week,
			"2024-08-08",
		},
	}
}
