package main

import(
	"github.com/gorilla/websocket"
)

//client represents a single chatting user
type client struct{
	// socket is the web socker this client
	socket	*websocket.Conn

	//send is a channel on which messages are sent
	send 	chan []byte

	//room is the room for this client is chatting on
	room	*room
}

func (c *client) read() {
	for{
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		} 
	}
	c.socket.Close()
}
