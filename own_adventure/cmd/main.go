package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gogojson/own_adventure/decoder"
	ownadventure "github.com/gogojson/own_adventure/own_adventure"
)

const layoutTemp = "layout.html"

type BookServer struct {
	BookData ownadventure.Book
}

func (b BookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(layoutTemp))
	if err := tmpl.Execute(w, b.BookData["intro"]); err != nil {
		panic(err)
	}

}

func main() {
	// Init
	filePath := flag.String("file", "gopher.json", "The file path used to run the story")
	port := flag.Int("port", 8080, "The port used to run the server")
	flag.Parse()

	fmt.Printf("Playing story with '%s' file\n", *filePath)
	book, err := decoder.JsonDecoder(*filePath)
	if err != nil {
		panic(err)
	}

	bookServer := BookServer{
		BookData: book,
	}

	// RUN server with the given flag
	fmt.Printf("Server Running in port :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), bookServer))

}

// {
// 	"intro": {
// 	  "title": "The Little Blue Gopher",
// 	  "story": [
// 		"Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
// 		"One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
// 		"On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit."
// 	  ],
// 	  "options": [
// 		{
// 		  "text": "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
// 		  "arc": "new-york"
// 		},
// 		{
// 		  "text": "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
// 		  "arc": "denver"
// 		}
// 	  ]
// 	},
