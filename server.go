package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("whisper"))
  http.Handle("/", fs)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}
