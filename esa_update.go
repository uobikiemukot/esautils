package main

import (
	"fmt"
	"github.com/uobikiemukot/go-esa/esa"
	"os"
)

func main() {
	apikey := os.Getenv("ESA_API_KEY")
	team := os.Getenv("ESA_TEAM")
	client := esa.NewClient(apikey)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ./esa_push PostDir")
		os.Exit(1)
	}

	if err := updatePost(client, team, os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "updatePost Failed: %v\n", err)
		os.Exit(2)
	}

	os.Exit(0)
}
