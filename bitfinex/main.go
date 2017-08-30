package main

import (
	"fmt"
	"gitlab.com/vitams/trade/bitfinex/lib/ws"
	//	"encoding/json"
	//	"gitlab.com/vitams/trade/bitfinex/lib/tools"
	// "reflect"
	"encoding/json"
)

func main() {

	// Init
	var s ws.Core

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

	type Tick struct {
		bid float64
	}

	for {
		var m map[string]interface{}
		msg := s.ReadMessage()
		fmt.Println(string(msg))
		json.Unmarshal(msg, &m)
		if value, ok := m["event"]; ok {
			switch value {
			case "subscribed":
				fmt.Println(m)
			}
		}
	}
}
