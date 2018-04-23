package main

import (
  "os"
  "github.com/upamune/go-esa/esa"
)

func main() {
  apikey := os.Getenv("ESA_API_KEY")
  team   := os.Getenv("ESA_TEAM")
  client := esa.NewClient(apikey)

  if len(os.Args) < 2 {
    panic("usage: ./esa_push PostDir")
  }

  update_post(client, team, os.Args[1])
}
