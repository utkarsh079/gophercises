package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/utkarsh079/gophercises/cyoa"
)

func main() {
	filename := flag.String("filename", "gopher.json", "name of the file that contains stories")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	story := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", story)
}
