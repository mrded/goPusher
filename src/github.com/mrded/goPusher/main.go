package main

import (
  "gopkg.in/antage/eventsource.v1"
  "github.com/gorilla/mux"
  
  "github.com/mrded/goPusher/cfg"

  "io/ioutil"
  "net/http"
  "gopkg.in/inconshreveable/log15.v2"
  "fmt"
)

func main() {
  options := cfg.GetOptions()

  log := log15.New()
  
  log.Info("Server is ready!", "port", options.Port, "token", options.Token)
  log.Info(fmt.Sprintf("Listening for post requests on http://localhost:%s/events", options.Port))
  log.Info(fmt.Sprintf("SSE streaming avaliable on http://localhost:%s/stream", options.Port))
  
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
      if token, ok := r.Header["X-Token"]; ok { 
        if token[0] == options.Token {
          event := mux.Vars(r)["event"]
          id := mux.Vars(r)["id"]
          
          data, err := ioutil.ReadAll(r.Body);
          if err != nil {
            log.Error("Cannot read body", "message", err)
          }
    
          es.SendEventMessage(string(data), event, id)
          log.Info("Message has been sent", "id", id, "event", event)
        
        } else {
          log.Error("The request has wrong token", "token", token[0])
          http.Error(w, "The request has wrong token", http.StatusUnauthorized)
        }
      } else {
        log.Error("The request doesn't contain authentication token")
        http.Error(w, "The request doesn't contain authentication token", http.StatusUnauthorized)
      } 
    } else {
      log.Error("Received wrong http request")
      http.Error(w, "POST requests only", http.StatusMethodNotAllowed)
    }
  })

  r.Handle("/", http.FileServer(http.Dir("./public")))
}
