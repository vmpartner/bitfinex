package main

import (
	"fmt"
	"encoding/json"
	"gitlab.com/vitams/trade/bitfinex/src"
)

func main() {

	// Init
	var b src.Bitfinex
	var s src.WebSocket

	// Connect & Reconnect
	URL := "wss://api.bitfinex.com/ws/2"
	s.HAConnect(URL)

	// Create message
	msg := make(map[string]string)
	msg["event"] = "subscribe"
	msg["channel"] = "ticker"
	msg["symbol"] = "tBTCUSD"

	// Send
	s.SendMessage(msg)

	// Auth
	auth := b.GetAuth()
	fmt.Println(auth)
	s.SendMessage(auth)

	type Tick struct {
		bid float64
	}

	for {
		var m map[string]interface{}
		msg := s.ReadMessage()
		json.Unmarshal(msg, &m)
		if value, ok := m["event"]; ok {
			switch value {
			case "subscribed":
			case "info":
			}
		}
		fmt.Println(string(msg))
	}
}
