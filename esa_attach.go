package main

import (
	"fmt"
	"os"
	"github.com/uobikiemukot/go-esa/esa"
)

func main() {
	apikey := os.Getenv("ESA_API_KEY")
	team   := os.Getenv("ESA_TEAM")
	client := esa.NewClient(apikey)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: ./esa_upload File")
		os.Exit(1)
	}

	url, err := client.Attachment.UploadAttachmentFile(team, os.Args[1])
	if err != nil {
		fmt.Printf("UploadAttachmentFile Failed: %v\n", err)
		os.Exit(2)
	}
	fmt.Println(url)
	os.Exit(0)
}
