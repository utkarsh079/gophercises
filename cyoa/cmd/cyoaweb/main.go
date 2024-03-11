package main

import (
	"encoding/json"
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

	var story cyoa.Story
	d := json.NewDecoder(f)
	if err := d.Decode(cyoa.Story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", story)
}
