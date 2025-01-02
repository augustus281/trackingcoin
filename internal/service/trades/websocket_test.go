package trades

import (
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

func TestConnect(t *testing.T) {
	serverURL := "wss://stream.coinmarketcap.com/ws"

	headers := http.Header{
		"Authorization":            []string{"Bearer 8ed074ba-2ab2-4306-ae28-52ad59c74fb2"},
		"User-Agent":               []string{"Golang-WebSocket-Client/1.0"},
		"Origin":                   []string{"https://coinmarketcap.com"},
		"Sec-WebSocket-Version":    []string{"13"},
		"Sec-WebSocket-Key":        []string{"CzZ7Va4UtKUQy989ndlNvg=="},
		"Sec-WebSocket-Extensions": []string{"permessage-deflate; client_max_window_bits"},
		"Host":                     []string{"stream.coinmarketcap.com"},
	}

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, headers)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to WebSocket")
}
