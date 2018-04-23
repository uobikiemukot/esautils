package main

import (
	"github.com/upamune/go-esa/esa"
	"os"
)

func main() {
	team := os.Getenv("ESA_TEAM")
	apikey := os.Getenv("ESA_API_KEY")
	client := esa.NewClient(apikey)

	if len(os.Args) < 3 {
		panic("usage: ./esa_create Category Name")
	}

	post := new_post(client, team, os.Args[1], os.Args[2])
	write_post(post)
}
