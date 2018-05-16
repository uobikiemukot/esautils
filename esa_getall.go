package main

import (
	"fmt"
	"github.com/uobikiemukot/go-esa/esa"
	"os"
)

func main() {
	team := os.Getenv("ESA_TEAM")
	user := os.Getenv("ESA_USER")
	apikey := os.Getenv("ESA_API_KEY")
	client := esa.NewClient(apikey)

	posts, err := getPosts(client, team, user)
	if err != err {
		fmt.Fprintf(os.Stderr, "getPosts Failed: %v\n", err)
		os.Exit(1)
	}

	if err = writePosts(posts); err != nil {
		fmt.Fprintf(os.Stderr, "writePosts Failed: %v\n", err)
		os.Exit(2)
	}

	os.Exit(0)
}
