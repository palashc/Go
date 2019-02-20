package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"flag"
	"strings"
)

//helps if input comes from something other than csv files
type problem struct {
	question string
	answer string
}

func main() {

	//get csv file through command line arguments
	var csvFileName = flag.String("csv", "problems.csv", "csv filename with problem set with format: question,answer")
	flag.Parse()

	//open and read file
	file, err := os.Open(*csvFileName)  //flag package returns pointer
	if err != nil {
		exit(fmt.Sprintf("Falied to open the csv file: %s", *csvFileName))
	}

	//handle csv
	r := csv.NewReader(file) // os.Open implements io.Reader interface
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse csv file")
	}
	problems := parseLines(lines)
	
	//show problems to user
	correct_answers := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		// read answer
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == problem.answer {
			correct_answers++
		}
	}
	fmt.Printf("You got %d/%d correct answers!", correct_answers, len(problems))
	
}

func parseLines(lines [][]string) []problem {
	 probs := make([]problem, len(lines))
	 for i, line := range lines {
	 	probs[i] = problem{
	 				question:line[0], 
	 				answer:strings.TrimSpace(line[1]),
	 				}
	 }
	 return probs
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}