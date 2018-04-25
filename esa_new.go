package main

import (
	"fmt"
	"github.com/upamune/go-esa/esa"
	"os"
)

func main() {
	team := os.Getenv("ESA_TEAM")
	apikey := os.Getenv("ESA_API_KEY")
	client := esa.NewClient(apikey)

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: ./esa_create Category Name")
		os.Exit(1)
	}

	post, err := newPost(client, team, os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "newPost Failed: %v\n", err)
		os.Exit(2)
	}

	if err = writePost(post); err != nil {
		fmt.Fprintf(os.Stderr, "writePost Failed: %v\n", err)
		os.Exit(3)
	}

	os.Exit(0)
}
