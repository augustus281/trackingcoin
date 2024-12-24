package dto

import "time"

type MarketPairsResponse struct {
	Data   DataMarketPairsResponse `json:"data"`
	Status Status                  `json:"status"`
}

type DataMarketPairsResponse struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	Symbol            string              `json:"symbol"`
	NumMarketPairs    int                 `json:"numMarketPairs"`
	MarketPairs       []MarketPair        `json:"marketPairs"`
	Sort              string              `json:"sort"`
	Direction         string              `json:"direction"`
	SponsoredExchange []SponsoredExchange `json:"sponsoredExchange"`
}

type QuoteLastestResponse struct {
	Data   []QuoteLastest `json:"data"`
	Status Status         `json:"status"`
}

type QuoteLastest struct {
	ID                            int64     `json:"id"`
	Name                          string    `json:"name"`
	Symbol                        string    `json:"symbol"`
	Slug                          string    `json:"slug"`
	Quotes                        []Quote   `json:"quotes"`
	Rank                          int64     `json:"rank"`
	DateAdded                     time.Time `json:"dateAdded"`
	MaxSupply                     float64   `json:"maxSupply"`
	TotalSupply                   float64   `json:"totalSupply"`
	SelfReportedCirculatingSupply float64   `json:"selfReportedCirculatingSupply"`
	CirculatingSupply             float64   `json:"circulatingSupply"`
	MarketCapPercentChange1h      float64   `json:"marketCapPercentChange1h"`
	MarketCapPercentChange        float64   `json:"marketCapPercentChange"`
	MarketCapPercentChange7d      float64   `json:"marketCapPercentChange7d"`
	MarketCapPercentChange30d     float64   `json:"marketCapPercentChange30d"`
	MarketCapPercentChange90d     float64   `json:"marketCapPercentChange90d"`
	MarketCapPercentChangeAll     float64   `json:"marketCapPercentChangeAll"`
	PercentChange1y               float64   `json:"percentChange1y"`
	PercentChange3y               float64   `json:"percentChange3y"`
	PercentChangeAll              float64   `json:"percentChangeAll"`
}

type Quote struct {
	Name                   string  `json:"name"`
	Price                  float64 `json:"price"`
	PercentChange1h        float64 `json:"percentChange1h"`
	PercentChange24h       float64 `json:"percentChange24h"`
	PercentChange7d        float64 `json:"percentChange7d"`
	PercentChange30d       float64 `json:"percentChange30d"`
	PercentChange60d       float64 `json:"percentChange60d"`
	PercentChange90d       float64 `json:"percentChange90d"`
	PercentChangeYtd       float64 `json:"percentChangeYtd"`
	MarketCap              float64 `json:"marketCap"`
	Volume24h              float64 `json:"volume24h"`
	FullyDilutedMarketCap  float64 `json:"fullyDilutedMarketCap"`
	MarketCapByTotalSupply float64 `json:"marketCapByTotalSupply"`
}

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    string    `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      string    `json:"elapsed"`
	CreditCount  int64     `json:"credit_count"`
}
