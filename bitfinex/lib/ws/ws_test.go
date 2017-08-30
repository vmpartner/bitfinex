package ws

import (
	"testing"
)

func TestCore_Connect(t *testing.T) {
	var core Core
	URL := "wss://api.bitfinex.com/ws/2"
	err := core.Connect(URL)
	if err != nil {
		t.Error(err)
	}
}

func TestCore_SendMessage(t *testing.T) {

}