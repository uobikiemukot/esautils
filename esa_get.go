package main

import (
	"github.com/upamune/go-esa/esa"
	"os"
	"strconv"
)

func main() {
	apikey := os.Getenv("ESA_API_KEY")
	team := os.Getenv("ESA_TEAM")
	client := esa.NewClient(apikey)

	if len(os.Args) < 2 {
		panic("usage: ./esa_get PostNumber")
	}

	post_num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("strconv.Atoi Failed" + err.Error())
	}

	post := getPost(client, team, post_num)
	writePost(post)
}
