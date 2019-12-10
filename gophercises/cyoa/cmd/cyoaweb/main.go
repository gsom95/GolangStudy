package main

import (
	"GolangStudy/gophercises/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "the port on which the server to start")
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()

	storyFile, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Can't open %s: %s\n", *filename, err)
	}

	story, err := cyoa.JSONToStory(storyFile)
	if err != nil {
		log.Fatalf("Can't decode JSON: %s\n", err)
	}

	h := cyoa.NewHandler(story)
	log.Printf("Starting the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
