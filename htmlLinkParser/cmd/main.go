package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var result = make(map[string]LinkEle)

func main() {

	r, err := os.ReadFile("testFiles/ex4.html")
	if err != nil {
		panic(err)
	}

	node, err := html.Parse(strings.NewReader(string(r)))
	if err != nil {
		panic(err)
	}

	htmlLooper(node)

	rList := make([]LinkEle, len(result))
	i := 0
	for _, v := range result {
		rList[i] = v
		i++
	}
	fmt.Printf("%+v\n", rList)
}

type LinkEle struct {
	Link string
	Text string
}

func htmlLooper(n *html.Node) {
	// fmt.Printf("Namespace %v\n", n.Namespace)
	// Check if it is an element
	if n.Type == html.ElementNode && len(n.Attr) != 0 {
		// Check if it has link
		for _, v := range n.Attr {
			if v.Key == "href" {
				result[v.Val] = LinkEle{Link: v.Val}
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					textFinder(c, v.Val)
				}
			}
		}
	}
	// if n.Type == html.TextNode
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		htmlLooper(c)
	}

}

func textFinder(n *html.Node, key string) {
	if n.Type == html.TextNode && n.Data != "" {
		temp := result[key]
		temp.Text += strings.TrimSpace(n.Data)
		result[key] = temp
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		textFinder(c, key)
	}

}
