package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	defaultFilePath  = "problems.csv"
	defaultTimeLimit = 30
)

func main() {
	// Setup flags
	var filePath = flag.String("f", defaultFilePath, "Change question file")
	var timeLimit = flag.Int("t", defaultTimeLimit, "Change the default time limit")
	flag.Parse()

	// Open the CSV file
	f, err := os.Open(*filePath)
	if err != nil {
		panic("Failed to open default csv file")
	}
	defer f.Close()
	reader := csv.NewReader(f)
	r, err := reader.ReadAll()
	if err != nil {
		panic("Fail to read csv file ")
	}

	// Setup scores
	var right = 0
	var total = len(r)

	startTime := time.Now() //TODO: It would be easier if we use the time.Timer
	// Iterate through questions
	ch := make(chan string)

	go func(ch chan string) {
		var a string
		if _, err := fmt.Scanln(&a); err != nil {
			panic(err)
		}
		ch <- a
	}(ch)

	go func(ch chan string) {
		for {
			if time.Since(startTime) > time.Second*time.Duration(*timeLimit) {
				ch <- "Timeout"
			}
		}
	}(ch)

	for i, v := range r {
		if len(v) < 2 {
			fmt.Println("Data issue")
			panic("File has data issue")
		}
		fmt.Printf("Q.%d %s= ", i+1, v[0])
		input := <-ch

		if input == "Timeout" {
			fScore := (right * 100) / total
			fmt.Println()
			fmt.Println("----------------------------------")
			fmt.Println("Time Out!")
			fmt.Printf("Got %d out of %d\n", right, total)
			fmt.Printf("Your final score is %d\n", fScore)
			return
		} else {
			if input == v[1] {
				right++
			}
		}

	}
	fScore := (right * 100) / total
	fmt.Println("----------------------------------")
	fmt.Println("Time Out!")
	fmt.Printf("Got %d out of %d\n", right, total)
	fmt.Printf("Your final score is %d\n", fScore)
}
