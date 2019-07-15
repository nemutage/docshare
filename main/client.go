package main

import (
  "github.com/gorilla/websocket"
)
 

type client struct {
  conn *websocket.Conn
  send chan []byte
  room *room
}


func (c *client) read() {
  for {
    if _, msg, err := c.conn.ReadMessage(); err == nil {
      c.room.forward <- msg
    } else {
      break
    }
  }
  c.conn.Close()
}


func (c *client) write() {
  for msg := range c.send {
    if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
      break
    }
  }
  c.conn.Close()
}