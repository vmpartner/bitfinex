package main

import (
	"fmt"
	"encoding/json"
	"gitlab.com/vitams/trade/bitfinex/src"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"reflect"
	"log"
)

func main() {

	// Init
	var b src.Bitfinex
	b.Channels = make(map[int]src.Channel)
	var s src.WebSocket

	// Open DB
	db, err := sql.Open("mysql", src.GetDSN())
	src.CheckErr(err)
	db.Ping()
	defer db.Close()

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
	s.SendMessage(auth)

	// Read messages
	for {

		// Init
		show := true
		var r interface{}
		msg := s.ReadMessage()
		json.Unmarshal(msg, &r)

		switch reflect.TypeOf(r).String() {

		case "map[string]interface {}": // Event data

			data := r.(map[string]interface{})
			if value, ok := data["event"]; ok {
				switch value {
				case "info":
					show = false
				case "auth":
					show = false
					if data["status"] != "OK" {
						panic("Error auth")
					}
				case "subscribed":
					show = false
					if data["channel"] == "ticker" {
						//fmt.Println(data)
						ch := src.Channel{
							Channel: data["channel"].(string),
							ChanId:  int(data["chanId"].(float64)),
						}
						b.Channels[int(data["chanId"].(float64))] = ch
					}
				}
			}

		case "[]interface {}": // Stream data

			data := r.([]interface{})
			channelId := int(data[0].(float64))

			// Skip heart bit
			if len(data) == 2 {
				if reflect.TypeOf(data[1]).String() == "string" {
					if data[1].(string) == "hb" {
						continue
					}
				}
			}

			// Get stream
			switch b.Channels[channelId].Channel {
			case "ticker":
				t := data[1].([]interface{})
				fmt.Println(t)
				ticker := src.Ticker{
					Bid:             t[0].(float64),
					BidSize:         t[1].(float64),
					Ask:             t[2].(float64),
					AskSize:         t[3].(float64),
					DailyChange:     t[4].(float64),
					DialyChangePerc: t[5].(float64),
					LastPrice:       t[6].(float64),
					Volume:          t[7].(float64),
					High:            t[8].(float64),
					Low:             t[9].(float64),
				}

				show = false
				log.Printf("Bitfinex Websocket Last %f Volume %f\n", ticker.LastPrice, ticker.Volume)

				_, err = db.Exec("INSERT INTO ticker(last_price, volume) VALUES (?, ?)", ticker.LastPrice, ticker.Volume)
				if err != nil {
					log.Fatal(err)
				}
			}

		default:
			fmt.Println(reflect.TypeOf(r).String())
		}

		if show {
			fmt.Println(string(msg))
		}
	}
}
