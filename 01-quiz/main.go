package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const ProblemsFile = "01-quiz/problems.csv"

var (
	correctAnswers int
	totalrecords   int
)

func main() {
	var (
		flagProblemsFilename = flag.String("p", ProblemsFile, "Problems file name")
		flagTimer            = flag.Duration("t", 30*time.Second, "Time of the test")
		flagShuffle          = flag.Bool("s", false, "Shuffle questions")
	)
	flag.Parse()

	// fmt.Println(flagTimer)
	if flagProblemsFilename == nil || flagTimer == nil || flagShuffle == nil {
		fmt.Println("Missing problems file name or timer")
		return
	}

	fmt.Println("Press Enter to start Quiz from %q in %v?\n", *flagProblemsFilename, *flagTimer)
	fmt.Scanln()
	f, err := os.Open(*flagProblemsFilename)
	if err != nil {
		fmt.Printf("Faild to open file: %v\n", err)
		return
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	totalrecords = len(records)
	if *flagShuffle {
		rand.Shuffle(totalrecords, func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}
	if err != nil {
		fmt.Printf("Falid to read csv file: %v\n", err)
		return
	}

	quizDone := startQuiz(records)

	quizTimer := time.Tick(*flagTimer)

	select {
	case <-quizDone:
	case <-quizTimer:
	}
	fmt.Printf("Result: %d/%d\n", correctAnswers, totalrecords)

}

func startQuiz(records [][]string) chan bool {
	done := make(chan bool)

	go func() {
		for i, record := range records {
			question, correctAnswer := record[0], record[1]
			fmt.Printf("%d. %s?\n", i+1, question)
			var answer string
			if _, err := fmt.Scan(&answer); err != nil {
				fmt.Printf("Faild to scan: %v\n", err)
				return
			}

			if answer == correctAnswer {
				correctAnswers++
			}
		}
		done <- true
	}()
	return done
}
