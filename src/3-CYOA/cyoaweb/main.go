package main

import (
	cyoa "3-CYOA/cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	story_file := flag.String("file", "gopher.json", "The json file with the story")
	port := flag.Int("port", 3000, "port number to start cyoa")
	flag.Parse()
	f, err := os.Open(*story_file)
	if err != nil {
		exit(fmt.Sprintf("Unable to open file: %s", story_file))
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		exit("Unable to decode json file")
	}

	h := cyoa.NewHandler(story, cyoa.WithPathFn(pathFn))
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("starting cyoa server at port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
