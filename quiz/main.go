package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	var csvfile string
	flag.StringVar(&csvfile, "csvFile", "./problems.csv", "Takes a file path to a csv")
	timeLimit := flag.Int("timeLimit", 30, "The time limit for the quizz in seconds.")
	flag.Parse()

	file, err := os.Open(csvfile)
	check(err)
	defer file.Close()
	var correctAnswers int
	lines, _ := csv.NewReader(file).ReadAll()

	fmt.Printf("Welcome to the quizz game, you have %v seconds to answer all the questions!\n", *timeLimit)
	fmt.Println("Enter your answer and press 'Enter'")

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	amountOfQuestions := len(lines)

	questionNum := 1
	for _, line := range lines {
		fmt.Printf("Question %v >  %v \n", questionNum, line[0])
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("Time is up!")
			fmt.Printf("You got %v/%v \n", correctAnswers, amountOfQuestions)
			return
		case answer := <-answerChannel:
			if answer == line[1] {
				fmt.Println("Correct Answer!")
				correctAnswers++
			} else {
				fmt.Println("Wrong Answer !")
			}
			questionNum++

		}
	}

	fmt.Printf("You got %v/%v \n", correctAnswers, amountOfQuestions)

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
