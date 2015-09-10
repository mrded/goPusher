package main

import (
  "gopkg.in/antage/eventsource.v1"
  "github.com/gorilla/mux"
  
  "github.com/mrded/goPusher/cfg"

  "io/ioutil"
  "net/http"
  "log"
  "fmt"
)

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
        []byte("Cache-Control: no-cache"),
      }
    },
  )

  defer es.Close()
  
  r := mux.NewRouter()
  
  r.Handle("/stream", es)

  r.HandleFunc("/event/{event}/{id}", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
      event := mux.Vars(r)["event"]
      id := mux.Vars(r)["id"]
      
      data, err := ioutil.ReadAll(r.Body);
      if err != nil {
        log.Fatal("Cannot read body; %s", err)
      }

      es.SendEventMessage(string(data), event, id)
      log.Printf("Message has been sent (id: %s, event: %s)", id, event)
    } else {
      http.Error(w, "POST only", http.StatusMethodNotAllowed)
    }
  })

  r.Handle("/", http.FileServer(http.Dir("./public")))
  
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", options.Port), r))
}
