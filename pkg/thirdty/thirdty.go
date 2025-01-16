package thirdty

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
)

func GetListingLatest(param dto.ListingParam) (*dto.ListingResponse, error) {
	response := &dto.ListingResponse{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		global.Logger.Error("failed to new request", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", global.Config.CoinMarket.APIKey)

	resp, err := client.Do(req)
	if err != nil {
		global.Logger.Error("error sending quest to server", zap.Error(err))
		return nil, err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(respBody, response); err != nil {
		global.Logger.Error("failed to unmarshal", zap.Error(err))
		return nil, err
	}

	return response, nil
}
