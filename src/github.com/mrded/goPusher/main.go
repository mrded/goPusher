package main

import (
  "gopkg.in/antage/eventsource.v1"
  "github.com/mrded/goPusher/cfg"

  "encoding/json"
  "net/http"
  "log"
  "fmt"
)

type Message struct {
  Id      string `json:"id"`
  Event   string `json:"event"`
  Data    string `json:"data"`
  Token   string `json:"token"`
}

func main() {
  options := cfg.GetOptions()

  log.Printf("Listening for post requests on http://localhost:%s/events", options.Port)
  log.Printf("SSE streaming avaliable on http://localhost:%s/stream", options.Port)

  log.Printf("Secret token is: %s", options.Token)

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

  http.Handle("/", http.FileServer(http.Dir("./public")))
  
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", options.Port), nil))
}
