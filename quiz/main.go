package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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

func RunQuiz(questionList []question, quizTime *time.Timer) int {
	correct := 0
	fmt.Println("The quiz is starting")
	for i, question := range questionList {
		ansChannel := make(chan string)
		fmt.Printf("Question number #%d: %s = ", i+1, question.q)
		go func() {
			reader := bufio.NewReader(os.Stdin)
			ans, _ := reader.ReadString('\n')
			ansChannel <- strings.TrimSpace(ans)
		}()

		select {
		case <-quizTime.C:
			fmt.Printf("\nTime is over")
			return correct
		case answer := <-ansChannel:
			if question.a == answer {
				correct++
			}
		}
	}
	return correct
}

func main() {
	csvFileName := flag.String("csv", "quiz.csv", "Name of the CSV")
	timeOut := flag.Int("timer", 20, "Total time for the quiz in seconds")
	flag.Parse()
	allQuestions, err := ReadCsv(*csvFileName)
	if err != nil {
		fmt.Print("unable to read csv file")
	}

	questionList := make([]question, len(allQuestions))

	for i, line := range allQuestions {
		questionList[i].q = line[0]
		questionList[i].a = line[1]
	}

	quizTimer := time.NewTimer(time.Duration(*timeOut) * time.Second)
	defer quizTimer.Stop()

	correct := RunQuiz(questionList, quizTimer)
	fmt.Printf("\nYou result is %d out of %d\n", correct, len(questionList))
}
