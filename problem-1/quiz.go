package gophercises

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	csv := flag.String("csv", "quiz.csv", "Name of the CSV")
	flag.Parse()
	csvName := *csv

	f, err := os.Open(csvName)
	if err != nil {
		log.Fatal(err)
	}

	allQuestions, err := csv.NewReader(f).ReadAll()
	if err != nil {
		Error("unable to read csv file")
	}
	fmt.Println(allQuestions)
}

func Error(msg string) {
	fmt.Errorf("Quiz failed with message: %s", msg)
}
