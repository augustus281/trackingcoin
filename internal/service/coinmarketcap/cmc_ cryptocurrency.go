package coinmarketcap

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/html"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
)

func (s *cmcService) GetDetailFromCMC(ctx *gin.Context, slug string) (*dto.Currency, error) {
	curRes := new(dto.Currency)
	url := global.Config.CoinMarket.CurrencyAPI
	resp, err := s.getHTML(ctx, url+strings.ToLower(slug))
	if err != nil {
		global.Logger.Error("error to request ", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		global.Logger.Error("parse body failed ", zap.Error(err))
		return nil, err
	}

	data := s.findData(node)
	if data == nil {
		global.Logger.Error("find data failed ", zap.Error(err))
		return nil, err
	}

	jsonTxt := strings.ReplaceAll(*data, "(MISSING)", "")
	err = json.Unmarshal([]byte(jsonTxt), &curRes)
	if err != nil {
		global.Logger.Error("failed to unmarshal json", zap.Error(err))
		return nil, err
	}

	return curRes, nil
}

func (r *cmcService) GetMarketPairFromCMC(ctx *gin.Context, slug string) (*dto.MarketPairsResponse, error) {
	marketpair := new(dto.MarketPairsResponse)
	url := global.Config.CoinMarket.URLApi + "cryptocurrency/market-pairs/latest?" + "slug=" + strings.ToLower(slug)

	resp, err := r.getHTML(ctx, url)
	if err != nil {
		global.Logger.Error("error to request ", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	byteResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error("error to read body ", zap.Error(err))
		return nil, err
	}

	err = json.Unmarshal([]byte(byteResp), &marketpair)
	if err != nil {
		global.Logger.Error("failed to unmarshal json", zap.Error(err))
		return nil, err
	}

	return marketpair, nil
}

func (s *cmcService) GetQuoteLastestFromCMC(ctx *gin.Context, id int) (*dto.QuoteLastestResponse, error) {
	quote := new(dto.QuoteLastestResponse)
	url := global.Config.CoinMarket.URLApi + "cryptocurrency/quote/latest?" + "id=" + strconv.Itoa(int(id))
	resp, err := s.getHTML(ctx, url)
	if err != nil {
		global.Logger.Error("error to request ", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	byteResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Logger.Error("error to read body ", zap.Error(err))
		return nil, err
	}

	err = json.Unmarshal([]byte(byteResp), &quote)
	if err != nil {
		global.Logger.Error("failed to unmarshal json", zap.Error(err))
		return nil, err
	}
	return quote, nil
}
