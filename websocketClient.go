package main

import (
	"github.com/gorilla/websocket"
)

func ConnectCore(CoreUrl string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial(CoreUrl, nil)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func SendWsMessage(SendMessage []byte, Conn *websocket.Conn) error {
	return Conn.WriteMessage(websocket.BinaryMessage, SendMessage)
}
