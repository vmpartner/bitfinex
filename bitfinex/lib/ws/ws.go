package ws

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"fmt"
	"time"
	"gitlab.com/vitams/trade/bitfinex/lib/tools"
)

type Core struct {
	Conn *websocket.Conn
}

func (c *Core) SendMessage(msg map[string]string) {
	b, err := json.Marshal(msg)
	tools.CheckErr(err)
	c.Conn.WriteMessage(websocket.TextMessage, []byte(string(b)))
}

func (c *Core) ReadMessage() []byte {
	_, msg, err := c.Conn.ReadMessage()
	tools.CheckErr(err)

	return msg
}

func (c *Core) Connect(URL string) (error) {
	var dialer *websocket.Dialer
	var err error
	c.Conn, _, err = dialer.Dial(URL, nil)

	return err
}

func (c *Core) HAConnect(URL string) {
	var err error
	for i := 0; i < 10; i++ {
		err = c.Connect(URL)
		if err != nil {
			fmt.Println("Error connect: ", err)
			fmt.Println("Try reconnect in 2 sec...")
			time.Sleep(time.Second * 2)
			continue
		}
		break
	}
}
