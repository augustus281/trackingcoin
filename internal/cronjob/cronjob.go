package cronjob

import (
	"context"
	"sync"
	"time"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
	"github.com/augustus281/trackingcoin/internal/util"
	"go.uber.org/zap"
)

func RunJob() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	emails, err := global.Db.Queries.GetAllEmails(ctx)
	if err != nil {
		global.Logger.Error("failed to get all emails", zap.Error(err))
		return err
	}

	// listingData, err := thirdty.GetListingLatest(dto.ListingParam{})
	// if err != nil {
	// 	global.Logger.Error("failed to get listing latest", zap.Error(err))
	// 	return err
	// }

	var wg sync.WaitGroup
	for _, email := range emails {
		wg.Add(1)
		go func() {
			listing := dto.EmailListing{
				Subject:       "Coin listing",
				RecipientName: email,
				ListingData: dto.ListingResponse{
					Data: []dto.CryptoData{
						{
							Name:        "Bitcoin",
							Symbol:      "BTC",
							LastUpdated: time.Now(),
							Quote: dto.QuoteCryptop{
								USD: dto.QuoteUSD{
									Price: 98830.11719584813,
								},
							},
						},
						{
							Name:        "Ethereum",
							Symbol:      "ETH",
							LastUpdated: time.Now(),
							Quote: dto.QuoteCryptop{
								USD: dto.QuoteUSD{
									Price: 3660.8926992061274,
								},
							},
						},
					},
				},
			}
			defer wg.Done()
			if err := util.SendMailHTML(email, listing); err != nil {
				global.Logger.Error("failed to send email", zap.String("email", email), zap.Error(err))
			}
		}()
	}

	wg.Wait()
	return nil
}

func StartHourlyJob() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		err := RunJob()
		if err != nil {
			global.Logger.Error("run job failed", zap.Error(err))
		}
	}
}
