package main

import (
  "github.com/luizbranco/hexagon"
  "github.com/codegangsta/negroni"
  "net/http"
  "encoding/json"
  "fmt"
)

func main() {
  mux := http.NewServeMux();

  hex := hexagon.NewHex(0,0)

  mux.HandleFunc("/map.json", func(w http.ResponseWriter, req *http.Request) {
    h, err := json.Marshal(hex)
    if err != nil {
      fmt.Printf("Error: %s", err)
      return
    }
    fmt.Fprintf(w, string(h))
  })

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(":3000")
}
