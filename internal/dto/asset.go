package dto

type MarketPair struct {
	ExchangeID          int64       `json:"exchangeId"`
	ExchangeName        string      `json:"exchangeName"`
	ExchangeSlug        string      `json:"exchangeSlug"`
	MarketID            int64       `json:"marketId"`
	Pair                string      `json:"marketPair"`
	Category            string      `json:"category"`
	MarketURL           string      `json:"marketUrl"`
	MarketReputation    float64     `json:"marketReputation"`
	BaseSymbol          string      `json:"baseSymbol"`
	QuoteSymbol         string      `json:"quoteSymbol"`
	QuoteCurrencyID     int64       `json:"quoteCurrencyId"`
	Price               float64     `json:"price"`
	VolumeUSD           float64     `json:"volumeUsd"`
	EffectiveLiquidity  float64     `json:"effectiveLiquidity"`
	Quote               float64     `json:"quote"`
	VolumeBase          float64     `json:"volumeBase"`
	VolumeQuote         float64     `json:"volumeQuote"`
	FeeType             string      `json:"feeType"`
	DepthUsdNegativeTwo float64     `json:"depthUsdNegativeTwo"`
	DepthUsdPositiveTwo float64     `json:"depthUsdPositiveTwo"`
	IsVerified          int64       `json:"isVerified"`
	Quotes              []QuoteData `json:"quotes"`
	Type                string      `json:"type"`
	CenterType          string      `json:"centerType"`
}

type QuoteData struct {
	ID               string  `json:"id"`
	Price            float64 `json:"price"`
	Volume24h        float64 `json:"volume24h"`
	DepthPositiveTwo float64 `json:"depthPositiveTwo"`
	DepthNegativeTwo float64 `json:"depthNegativeTwo"`
	IndexPrice       float64 `json:"indexPrice"`
}

type SponsoredExchange struct {
	EventID       int64     `json:"eventId"`
	CustomName    string    `json:"customName"`
	CustomLogo    string    `json:"customLogo"`
	SubInfos      []SubInfo `json:"subInfos"`
	CustomOptions string    `json:"customOptions"`
}

type SubInfo struct {
	URL              string   `json:"url"`
	ShowMsg          string   `json:"showMsg"`
	SupportCountries []string `json:"supportCountries"`
	ExcludeCountries []string `json:"excludeCountries"`
}
