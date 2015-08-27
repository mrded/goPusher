package main

import (
  "gopkg.in/antage/eventsource.v1"

  "encoding/json"
  "net/http"
  "log"
)

type Message struct {
  Id      string `json:"id"`
  Event   string `json:"event"`
  Data    string `json:"data"`
}

func main() {
  es := eventsource.New(
    eventsource.DefaultSettings(),
    func(req *http.Request) [][]byte {
      return [][]byte{
        []byte("X-Accel-Buffering: no"),
        []byte("Access-Control-Allow-Origin: *"),
      }
    },
  )

  defer es.Close()
  http.Handle("/stream", es)

  http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
      var message Message

      decoder := json.NewDecoder(r.Body)

      // Read post message.
      if err := decoder.Decode(&message); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      es.SendEventMessage(message.Data, message.Event, message.Id)
      log.Printf("Message has been sent (id: %s, event: %s)", message.Id, message.Event)
    } else {
      http.Error(w, "POST only", http.StatusMethodNotAllowed)
    }
  })

  log.Fatal(http.ListenAndServe(":9090", nil))
}
