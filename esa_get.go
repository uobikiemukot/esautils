package main

import (
	"fmt"
	"github.com/uobikiemukot/go-esa/esa"
	"os"
	"strconv"
)

func main() {
	apikey := os.Getenv("ESA_API_KEY")
	team := os.Getenv("ESA_TEAM")
	client := esa.NewClient(apikey)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ./esa_get PostNumber\n")
		os.Exit(1)
	}

	post_num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "strconv.Atoi Failed: %v\n", err)
		os.Exit(2)
	}

	post, err := getPost(client, team, post_num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getPost Failed: %v\n", err)
		os.Exit(3)
	}

	if err = writePost(post); err != nil {
		fmt.Fprintf(os.Stderr, "writePost Failed: %v\n", err)
		os.Exit(4)
	}

	os.Exit(0)
}
