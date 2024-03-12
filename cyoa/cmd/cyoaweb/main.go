package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"cyoa"
)

func main() {
	port := flag.Int("port", 3030, "port to run the cyo on")
	filename := flag.String("filename", "gopher.json", "name of the file that contains stories")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("starting a server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
