package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gogojson/own_adventure/decoder"
	ownadventure "github.com/gogojson/own_adventure/own_adventure"
)

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

	tmplOpts := ownadventure.WithTmpl(template.Must(template.ParseFiles("layout2.html")))

	bServer, err := ownadventure.NewBookServer(book, tmplOpts, ownadventure.WithPathFunc(pathFn))
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/story/", bServer)

	bServer2, err := ownadventure.NewBookServer(book)
	if err != nil {
		panic(err)
	}
	mux.Handle("/", bServer2)

	// RUN server with the given flag
	fmt.Printf("Server Running in port :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func pathFn(r *http.Request) string {
	base := "/story"
	url := r.URL.Path

	if len(url) < len(base)+2 || url == base || url == fmt.Sprintf("%s/", base) {
		url = "/story/intro"
	}

	url = url[len(base)+1:]
	fmt.Println(url)
	return url
}
