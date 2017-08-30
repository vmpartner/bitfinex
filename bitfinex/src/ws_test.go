package src

import (
	"testing"
)

func TestCore_Connect(t *testing.T) {
	var w WebSocket
	URL := "wss://api.bitfinex.com/ws/2"
	err := w.Connect(URL)
	if err != nil {
		t.Error(err)
	}
}

func TestCore_SendMessage(t *testing.T) {

}