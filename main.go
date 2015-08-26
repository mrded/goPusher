package main

import (
  "fmt"
  "encoding/json"
  "net/http"
)

type Message struct {
  Name    string `json:"name"`
  Channel string `json:"channel"`
  Data    string `json:"data"`
  Token   string `json:"token"`
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
  })

  http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
      var message Message

      decoder := json.NewDecoder(r.Body)

      // Read post message.
      if err := decoder.Decode(&message); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      // Converts message back to json.
      output, err := json.Marshal(message);
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      w.Header().Set("Content-Type", "application/json")
      w.Write(output)
    } else {
      http.Error(w, "POST only", http.StatusMethodNotAllowed)
    }
  })

  http.ListenAndServe(":9090", nil)
}
