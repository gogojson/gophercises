package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	r, err := os.ReadFile("testFiles/ex1.html")
	if err != nil {
		panic(err)
	}

	node, err := html.Parse(strings.NewReader(string(r)))
	if err != nil {
		panic(err)
	}

	htmlLooper(node)
}

type LinkEle struct {
	Link string
	Text string
}

func htmlLooper(n *html.Node) {

	fmt.Println("---------------------------")
	fmt.Printf("Type %v\n", n.Type)
	fmt.Printf("Attr %v\n", n.Attr)
	fmt.Printf("Data %v\n", n.Data)
	fmt.Printf("Namespace %v\n", n.Namespace)
	// if n.Type == html.ElementNode {
	// 	fmt.Printf("Attr %v\n", n.Attr)
	// 	fmt.Printf("Data %v\n", n.Data)
	// 	fmt.Printf("Namespace %v\n", n.Namespace)

	// }
	// if n.Type == html.TextNode
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		htmlLooper(c)
	}

}
