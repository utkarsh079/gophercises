package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type question struct {
	q string
	a string
}

func ReadCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read csv file")
	}
	defer f.Close()

	allQuestions, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return nil, fmt.Errorf("unable to read csv file")
	}

	return allQuestions, nil
}

func RunQuiz(questionList []question) {
	correct := 0
	fmt.Println("The quiz is starting")
	for i, question := range questionList {
		var ans string
		fmt.Printf("Question number #%d: %s\n", i+1, question.q)
		fmt.Scanf("%s", &ans)
		if question.a == strings.TrimSpace(ans) {
			correct++
		}
	}
	fmt.Printf("You result is %d out of %d\n", correct, len(questionList))
}

func main() {
	filename := flag.String("csv", "quiz.csv", "Name of the CSV")
	flag.Parse()
	allQuestions, err := ReadCsv(*filename)
	if err != nil {
		fmt.Print("unable to read csv file")
	}

	questionList := make([]question, len(allQuestions))

	for i, line := range allQuestions {
		questionList[i].q = line[0]
		questionList[i].a = line[1]
	}

	RunQuiz(questionList)
}
