package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := flag.String("url", "https://courses.calhoun.io/", "URL to create an sitemap")
	flag.Parse()
	fmt.Println(*url)

	resp, err := http.Get(*url)
	if err != nil {
		panic("Invalid url")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Can not read resp body")
	}
	fmt.Println(string(b))
	httpLinkParser.init
}
