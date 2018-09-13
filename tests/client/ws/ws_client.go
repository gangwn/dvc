package ws

import (
				"github.com/gorilla/websocket"
	)

type WSClientHandler = func(message []byte)


type WSClient struct {
	Socket *websocket.Conn
	Send   chan []byte
	Handler WSClientHandler
}


func (c *WSClient) Read() {

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			c.Socket.Close()
			break
		}
		c.Handler(message)
	}
}

func (c *WSClient) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}