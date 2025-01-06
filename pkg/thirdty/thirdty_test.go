package thirdty

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/augustus281/trackingcoin/internal/dto"
)

func TestListing(t *testing.T) {
	response := &dto.ListingResponse{}
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		t.Fatalf("failed to create new request: %v", err)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "8ed074ba-2ab2-4306-ae28-52ad59c74fb2")

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("error sending request to server: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	if err := json.Unmarshal(respBody, response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	fmt.Println(response.Data[0])
}
