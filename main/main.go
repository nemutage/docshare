package main

import (
  "html/template"
  "log"
  "net/http"
)


var filename = "html/chat2.html"


func main() {
  r := &room{
    forward: make(chan []byte),
    join: make(chan *client),
    leave: make(chan *client),
    clients: make(map[*client]bool),
  }

  http.HandleFunc("/", serveHTML)
  http.Handle("/room", r)
  go r.run()
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}


func serveHTML(w http.ResponseWriter, r *http.Request) {
  t := template.Must(template.ParseFiles(filename))
  t.Execute(w, r)
}
