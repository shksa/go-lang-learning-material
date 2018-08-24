package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

const defaultFilePath = "/Users/sreekarnimbalkar/go/src/github.com/shksa/learningGo/gophersizes/quizGame/questions.csv"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// flag.String() defines a string flag with specified name, default value, and usage string
	// returns a pointer to the value of the flag.
	filenamePtr := flag.String("filename", defaultFilePath, "filename")
	// flag.Parse() parses the command-line flags from os.Args[1:].
	// Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	// fmt.Printf("filenamePtr: %s \n", *filenamePtr)
	// os.Open() opens the named file for reading.
	// methods on the returned file can be used for reading.
	// The returned file satisfies the io.Reader interface i.e has a method with signature Read func([]byte) (int, error)
	file, err := os.Open(*filenamePtr)
	check(err)
	// bufio.NewReader() takes in a io.Reader and returns a new io.Reader that has methods which support buffered reading.
	// csv.NewReade() takes in a io.Reader and returns a new io.Reader that has methods which support reading a csv file content
	csvReader := csv.NewReader(bufio.NewReader(file))
	// csvReader.ReadAll() reads all the remaining records from csvReader.
	// Each record is a slice of fields.
	questionAnswerRecords, err := csvReader.ReadAll()
	// fmt.Printf("%#v \n", questionAnswerRecords)
	var question, answer, answerEntered string
	var score int
	for qNum, questionAnswerPair := range questionAnswerRecords {
		question, answer = questionAnswerPair[0], questionAnswerPair[1]
		fmt.Printf("\nQuestion %d: %s \nAnswer %d: ", qNum, question, qNum)
		fmt.Scanf("%s", &answerEntered)
		if answerEntered == answer {
			score++
		}
	}
	fmt.Printf("\nTotal no. questions: %d \nNo. of questions answered correctly: %d\n", len(questionAnswerRecords), score)
}
