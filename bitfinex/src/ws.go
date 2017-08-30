package src

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"fmt"
	"time"
)

type WebSocket struct {
	Conn *websocket.Conn
}

func (s *WebSocket) SendMessage(msg interface{}) {
	b, err := json.Marshal(msg)
	CheckErr(err)
	s.Conn.WriteMessage(websocket.TextMessage, []byte(string(b)))
}

func (s *WebSocket) ReadMessage() []byte {
	_, msg, err := s.Conn.ReadMessage()
	CheckErr(err)

	return msg
}

func (s *WebSocket) Connect(URL string) (error) {
	var dialer *websocket.Dialer
	var err error
	s.Conn, _, err = dialer.Dial(URL, nil)

	return err
}

func (s *WebSocket) HAConnect(URL string) {
	var err error
	for i := 0; i < 10; i++ {
		err = s.Connect(URL)
		if err != nil {
			fmt.Println("Error connect: ", err)
			fmt.Println("Try reconnect in 2 sec...")
			time.Sleep(time.Second * 2)
			continue
		}
		break
	}
}
