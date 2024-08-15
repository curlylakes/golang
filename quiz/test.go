// quiz game
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	var file = "problems.csv"
	var data = readCsv(file)
	var questions = questions(data)
	var answers = answers(data)
	var runtime = len(questions)
	var userAnswers []string
	for i := 0; i < runtime; i++ {
		fmt.Println(questions[i])
		userAnswers = userInput(userAnswers)
	}
	var result = gameLogic(answers, userAnswers)

	fmt.Println("You got", result, "questions right!")

	fmt.Println("Your answers")
	fmt.Println(userAnswers)

	fmt.Println("The right answers")
	fmt.Println(answers)
}

func userInput(userAnswers []string) []string {
	var input string
	fmt.Print("Ans: ")
	fmt.Scan(&input)
	userAnswers = append(userAnswers, input)
	return userAnswers
}

func readCsv(filePath string) [][]string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Error whilst reading", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading file contet")
	}
	return records
}

func questions(file [][]string) []string {
	var questions []string
	for eachQuestion := range file {
		questions = append(questions, file[eachQuestion][0])
	}
	return questions
}

func answers(file [][]string) []string {
	var answers []string
	for eachAnswer := range file {
		answers = append(answers, file[eachAnswer][1])
	}
	return answers
}

func gameLogic(answers []string, userAnswers []string) int {
	var counter int
	for i := range len(answers) {
		if answers[i] == userAnswers[i] {
			counter++
		} else {
			counter = counter + 0
		}
	}
	return counter
}
