package main

import (
	"github.com/upamune/go-esa/esa"
	"os"
)

func main() {
	apikey := os.Getenv("ESA_API_KEY")
	team := os.Getenv("ESA_TEAM")
	client := esa.NewClient(apikey)

	if len(os.Args) < 2 {
		panic("usage: ./esa_push PostDir")
	}

	updatePost(client, team, os.Args[1])
}
