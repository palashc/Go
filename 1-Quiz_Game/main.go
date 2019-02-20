package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//helps if input comes from something other than csv files
type problem struct {
	question string
	answer   string
}

func main() {

	//get csv file through command line arguments
	var csvFileName = flag.String("csv", "problems.csv", "csv filename with problem set with format: question,answer")
	timeLimit := flag.Int("limit", 13, "time limit (sec) for the quiz")
	flag.Parse()

	//open and read file
	file, err := os.Open(*csvFileName) //flag package returns pointer
	if err != nil {
		exit(fmt.Sprintf("Falied to open the csv file: %s", *csvFileName))
	}

	//handle csv
	r := csv.NewReader(file) // os.Open implements io.Reader interface
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse csv file")
	}
	problemsMap := parseLines(lines)
	//timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//show problems to user
	correct_answers := 0
	qNum := 0
	for q, a := range problemsMap {
		qNum++
		fmt.Printf("Problem #%d: %s = ", qNum, q)
		answerChannel := make(chan string)
		go func() {
			// read answer
			//use goroutine to handle blocking of scanf
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerChannel <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou got %d/%d correct answers!", correct_answers, len(problemsMap))
			return
		case ans := <-answerChannel:
			if ans == a {
				correct_answers++
			}
		}
	}
	fmt.Printf("You got %d/%d correct answers!", correct_answers, len(problemsMap))
}

func parseLines(lines [][]string) map[string]string {
	//Use map instead of struct array since iteration over maps in Go is random
	qaMap := make(map[string]string, len(lines))
	for _, line := range lines {
		qaMap[line[0]] = strings.TrimSpace(line[1])
	}
	return qaMap
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
