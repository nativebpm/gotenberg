package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nativebpm/gotenberg"
)

func main() {
	client, err := gotenberg.NewClient(&http.Client{}, "http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	html := strings.NewReader("<html><body><h1>Hello World!</h1></body></html>")

	resp, err := client.Chromium().
		ConvertHTML(context.Background(), html).
		PaperSizeA6().
		Margins(0.5, 0.5, 0.5, 0.5).
		Send()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	f, _ := os.Create("out.pdf")
	io.Copy(f, resp.Body)
	f.Close()
}
