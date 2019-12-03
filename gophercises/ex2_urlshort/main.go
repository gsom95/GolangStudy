package main

import (
	"GolangStudy/gophercises/ex2_urlshort/urlshort"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	mapType := flag.String("type", "yaml", "type of path to url mapping")
	filename := flag.String("file", "map.yaml", "filename of map file")
	flag.Parse()

	mapFile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Can't open map file:", err)
	}
	defer mapFile.Close()

	fileContent, err := ioutil.ReadAll(mapFile)
	if err != nil {
		log.Fatalln("Can't read file:", err)
	}

	mux := defaultMux()
	// Build the MapHandler using the mux as the fallback
	defaultPathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(defaultPathsToUrls, mux)

	var handler http.HandlerFunc
	// Build the Handler using the mapHandler as the fallback
	handler, err = urlshort.CreateHandler([]byte(fileContent), *mapType, mapHandler)
	if err != nil {
		log.Fatalln("Can't create a handler:", err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
