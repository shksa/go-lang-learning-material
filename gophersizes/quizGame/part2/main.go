package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

const defaultFilePath = "/Users/sreekarnimbalkar/go/src/github.com/shksa/learningGo/gophersizes/quizGame/questions.csv"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseCLIFlags() (*string, *int) {
	filename := flag.String("filename", defaultFilePath, "filename of the quiz content")
	timeout := flag.Int("timeout", 30, "timeout for the quiz")
	flag.Parse()
	return filename, timeout
}

func parseCSVFile(filename string) [][]string {
	file, err1 := os.Open(filename)
	check(err1)
	csvReader := csv.NewReader(bufio.NewReader(file))
	questionAnswerRecords, err2 := csvReader.ReadAll()
	check(err2)
	return questionAnswerRecords
}

func startQuiz(questionAnswerRecords [][]string, scoreCh chan int) {
	var question, answer, answerEntered string
	var score int
	for qNum, questionAnswerPair := range questionAnswerRecords {
		question, answer = questionAnswerPair[0], questionAnswerPair[1]
		fmt.Printf("\nQuestion %d: %s \nAnswer %d: ", qNum, question, qNum)
		fmt.Scanf("%s", &answerEntered)
		if answerEntered == answer {
			<-scoreCh
			score++
			scoreCh <- score
		}
	}
}

func startTimer(timeout int, timeOverCh chan bool) {
	time.Sleep(time.Duration(timeout) * time.Second)
	timeOverCh <- true
}

func main() {
	filename, timeout := parseCLIFlags()
	questionAnswerRecords := parseCSVFile(*filename)
	fmt.Printf("\nPress enter to start the quiz: ")
	fmt.Scanf("%s")
	scoreCh := make(chan int, 1)
	scoreCh <- 0
	timeOverCh := make(chan bool)
	go startTimer(*timeout, timeOverCh)
	go startQuiz(questionAnswerRecords, scoreCh)
	<-timeOverCh // main goroutine is blocked here untill startTimer goroutine sends a value to the timeOverCh, indicating the timer has expired.
	fmt.Printf("\n\nTime out!\nTotal no. questions: %d \nNo. of questions answered correctly: %d\n\n", len(questionAnswerRecords), <-scoreCh)
}
