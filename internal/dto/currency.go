package dto

type Currency struct {
	Props                 Props       `json:"props"`
	Page                  string      `json:"page"`
	Query                 interface{} `json:"query"`
	BuildId               string      `json:"buildId"`
	AssetPrefix           string      `json:"assetPrefix"`
	IsFallback            bool        `json:"isFallback"`
	IsExperimentalCompile bool        `json:"isExperimentalCompile"`
	Gip                   bool        `json:"gip"`
	AppGip                bool        `json:"appGip"`
	ScriptLoader          interface{} `json:"scriptLoader"`
}

type Props struct {
	InitialI18nStore   interface{} `json:"initialI18nStore"`
	InitialLanguage    string      `json:"initialLanguage"`
	I18nServerInstance interface{} `json:"i18nServerInstance"`
	PageProps          PageProps   `json:"pageProps"`
}

type PageProps struct {
	DetailRes DetailRes `json:"detailRes"`
}

type DetailRes struct {
	Detail          Detail      `json:"detail"`
	Trending        interface{} `json:"trending"`
	GravityFlag     bool        `json:"gravityFlag"`
	ProjectInfoFlag bool        `json:"projectInfoFlag"`
}

type Detail struct {
	Id                        int         `json:"id"`
	Name                      string      `json:"name"`
	Symbol                    string      `json:"symbol"`
	Slug                      string      `json:"slug"`
	Category                  string      `json:"category"`
	Status                    string      `json:"status"`
	SubStatus                 string      `json:"subStatus"`
	Notice                    string      `json:"notice,omitempty"`
	AlertType                 int         `json:"alertType,omitempty"`
	AlertLink                 string      `json:"alertLink,omitempty"`
	WatchCount                string      `json:"watchCount"`
	WatchListRanking          int         `json:"watchListRanking"`
	LatestAdded               bool        `json:"latestAdded"`
	LaunchPrice               float64     `json:"launchPrice"`
	Tags                      []Tag       `json:"tags"`
	Urls                      Urls        `json:"urls"`
	Volume                    float64     `json:"volume"`
	VolumeChangePercentage24h float64     `json:"volumeChangePercentage24h"`
	CexVolume                 float64     `json:"cexVolume"`
	DexVolume                 float64     `json:"dexVolume"`
	Statistics                Statistics  `json:"statistics"`
	Quotes                    interface{} `json:"quotes"`
	Wallets                   interface{} `json:"wallets"`
	IsAudited                 bool        `json:"isAudited"`
	AuditInfos                interface{} `json:"auditInfos"`
	Holders                   interface{} `json:"holders"`
	DisplayTV                 int         `json:"displayTV"`
	IsInfiniteMaxSupply       int         `json:"isInfiniteMaxSupply"`
	TvCoinSymbol              string      `json:"tvCoinSymbol"`
	SupportWalletInfos        interface{} `json:"supportWalletInfos"`
	HoldersFlag               bool        `json:"holdersFlag"`
	RatingsFlag               bool        `json:"ratingsFlag"`
}

type Tag struct {
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Urls struct {
	Website      []string `json:"website"`
	TechnicalDoc []string `json:"technical_doc"`
	Explorer     []string `json:"explorer"`
	SourceCode   []string `json:"source_code"`
	MessageBoard []string `json:"message_board"`
	Chat         []string `json:"chat,omitempty"`
	Announcement []string `json:"announcement,omitempty"`
	Reddit       []string `json:"reddit"`
	Facebook     []string `json:"facebook,omitempty"`
	Twitter      []string `json:"twitter,omitempty"`
}

type Statistics struct {
	Price                                    float64 `json:"price"`
	PriceChangePercentage1h                  float64 `json:"priceChangePercentage1h"`
	PriceChangePercentage24h                 float64 `json:"priceChangePercentage24h"`
	PriceChangePercentage7d                  float64 `json:"priceChangePercentage7d"`
	PriceChangePercentage30d                 float64 `json:"priceChangePercentage30d"`
	PriceChangePercentage60d                 float64 `json:"priceChangePercentage60d"`
	PriceChangePercentage90d                 float64 `json:"priceChangePercentage90d"`
	PriceChangePercentage1y                  float64 `json:"priceChangePercentage1y"`
	PriceChangePercentageAll                 float64 `json:"priceChangePercentageAll"`
	MarketCap                                float64 `json:"marketCap"`
	MarketCapChangePercentage24h             float64 `json:"marketCapChangePercentage24h"`
	FullyDilutedMarketCap                    float64 `json:"fullyDilutedMarketCap"`
	FullyDilutedMarketCapChangePercentage24h float64 `json:"fullyDilutedMarketCapChangePercentage24h"`
	CirculatingSupply                        float64 `json:"circulatingSupply"`
	TotalSupply                              float64 `json:"totalSupply"`
	MaxSupply                                float64 `json:"maxSupply"`
	MarketCapDominance                       float64 `json:"marketCapDominance"`
	Rank                                     int     `json:"rank"`
	ROI                                      float64 `json:"roi"`
	Low24h                                   float64 `json:"low24h"`
	High24h                                  float64 `json:"high24h"`
	Low7d                                    float64 `json:"low7d"`
	High7d                                   float64 `json:"high7d"`
	Low30d                                   float64 `json:"low30d"`
	High30d                                  float64 `json:"high30d"`
	Low90d                                   float64 `json:"low90d"`
	High90d                                  float64 `json:"high90d"`
	Low52w                                   float64 `json:"low52w"`
	High52w                                  float64 `json:"high52w"`
	LowAllTime                               float64 `json:"lowAllTime"`
	HighAllTime                              float64 `json:"highAllTime"`
	LowAllTimeChangePercentage               float64 `json:"lowAllTimeChangePercentage"`
	HighAllTimeChangePercentage              float64 `json:"highAllTimeChangePercentage"`
	LowAllTimeTimestamp                      string  `json:"lowAllTimeTimestamp"`
	HighAllTimeTimestamp                     string  `json:"highAllTimeTimestamp"`
	LowYesterday                             float64 `json:"lowYesterday"`
	HighYesterday                            float64 `json:"highYesterday"`
	OpenYesterday                            float64 `json:"openYesterday"`
	CloseYesterday                           float64 `json:"closeYesterday"`
	PriceChangePercentageYesterday           float64 `json:"priceChangePercentageYesterday"`
	VolumeYesterday                          float64 `json:"volumeYesterday"`
	Turnover                                 float64 `json:"turnover"`
	YTDPriceChangePercentage                 float64 `json:"ytdPriceChangePercentage"`
	VolumeRank                               int     `json:"volumeRank"`
	VolumeMcRank                             int     `json:"volumeMcRank"`
	McTotalNum                               int     `json:"mcTotalNum"`
	VolumeTotalNum                           int     `json:"volumeTotalNum"`
	VolumeMcTotalNum                         int     `json:"volumeMcTotalNum"`
	Status                                   string  `json:"status"`
}
