package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		links := extractLinks(resp.Body)
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func extractLinks(r io.Reader) []string {
	links := make([]string, 0)
	tokenizer := html.NewTokenizer(r)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						if !strings.HasPrefix(attr.Val, "#") {
							links = append(links, attr.Val)
						}
					}
				}
			}
		}
	}
}
