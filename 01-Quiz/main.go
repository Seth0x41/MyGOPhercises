package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const problemsFile = "problems.csv"

func main() {
	// open CSV file and handling errors
	f, err := os.Open(problemsFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer f.Close()
	// read CSV file
	r := csv.NewReader(f)
	questions, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Error To Read File : %v\n", err)
		return
	}
	// Split and Display Questions for answers
	var correctAnswers int
	var answer string
	for i, record := range questions {
		question, correctanswer := record[0], record[1]
		fmt.Printf("%d- %s\n", i+1, question)
		fmt.Scan(&answer)
		if err != nil {
			fmt.Printf("Faild To Scan %v\n", err)
			return
		}
		// Check if answer is correct
		if answer == correctanswer {
			correctAnswers++
		}
		// The Result
	}
	fmt.Printf(" Result %d / %d", correctAnswers, len(questions))

}
