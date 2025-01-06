package dto

import "time"

type EmailListing struct {
	Subject       string
	RecipientName string
	ListingData   ListingResponse
}

type ListingResponse struct {
	Status ListingStatus `json:"status"`
	Data   []CryptoData  `json:"data"`
}

type CryptoData struct {
	ID                            int          `json:"id"`
	Name                          string       `json:"name"`
	Symbol                        string       `json:"symbol"`
	Slug                          string       `json:"slug"`
	NumMarketPairs                int          `json:"num_market_pairs"`
	DateAdded                     time.Time    `json:"date_added"`
	Tags                          []string     `json:"tags"`
	MaxSupply                     float64      `json:"max_supply"`
	CirculatingSupply             float64      `json:"circulating_supply"`
	TotalSupply                   float64      `json:"total_supply"`
	InfiniteSupply                bool         `json:"infinite_supply"`
	Platform                      PlatForm     `json:"platform"`
	CMCRank                       int          `json:"cmc_rank"`
	SelfReportedCirculatingSupply float64      `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         float64      `json:"self_reported_market_cap"`
	TVLRatio                      float64      `json:"tvl_ratio"`
	LastUpdated                   time.Time    `json:"last_updated"`
	Quote                         QuoteCryptop `json:"quote"`
	Price                         float64      `json:"price"`
}

type PlatForm struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Slug         string `json:"slug"`
	TokenAddress string `json:"token_address"`
}

type QuoteUSD struct {
	Price                 float64   `json:"price"`
	Volume24h             float64   `json:"volume_24h"`
	VolumeChange24h       float64   `json:"volume_change_24h"`
	PercentChange1h       float64   `json:"percent_change_1h"`
	PercentChange24h      float64   `json:"percent_change_24h"`
	PercentChange7d       float64   `json:"percent_change_7d"`
	PercentChange30d      float64   `json:"percent_change_30d"`
	PercentChange60d      float64   `json:"percent_change_60d"`
	PercentChange90d      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	TVL                   *float64  `json:"tvl"`
	LastUpdated           time.Time `json:"last_updated"`
}

type QuoteCryptop struct {
	USD QuoteUSD `json:"USD"`
}

type ListingStatus struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
	Notice       string `json:"notice"`
	TotalCount   int    `json:"total_count"`
}
