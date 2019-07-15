package main

import (
  "net/http"
  "log"
  "github.com/gorilla/websocket"
)


type room struct {
  forward chan []byte
  join chan *client
  leave chan *client
  clients map[*client]bool
}


func (r *room) run() {
  for {
    select {
    case client := <-r.join:
      r.clients[client] = true
    case client := <-r.leave:
      delete(r.clients, client)
      close(client.send)
    case msg := <-r.forward:
      for client := range r.clients {
        select {
        case client.send <- msg:
        default:
          delete(r.clients, client)
          close(client.send)
        }
      }
    }
  }
}


var upgrader = &websocket.Upgrader{
  ReadBufferSize: 1024, 
  WriteBufferSize: 1024,
}


func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  conn, err := upgrader.Upgrade(w, req, nil)
  if err != nil {
    log.Println(err)
    return
  }

  cli := &client{
    conn: conn,
    send: make(chan []byte, 1024),
    room: r,
  }

  r.join <- cli
  defer func() {
    r.leave <- cli
  }()
  go cli.write()
  cli.read()
}






