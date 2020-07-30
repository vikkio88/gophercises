package lib

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Quiz Struct that holds a single question of the quiz
type Quiz struct {
	ID            int
	Question      string
	Answer        string
	CorrectAnswer string
}

func (q *Quiz) registerAnswer(answer string) bool {
	q.Answer = answer
	return q.Answer == q.CorrectAnswer
}

// LoadQuiz function that loads a csv file and returns a slice of Quiz
func LoadQuiz(filePath string) []Quiz {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file %s\n", filePath)
		os.Exit(1)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(records)

	//result := make([]Quiz, 0)
	// we know how bit it is no need to leave it to HEAP
	result := make([]Quiz, len(records))

	for i, row := range records {
		quiz := Quiz{ID: i, Question: row[0], CorrectAnswer: strings.TrimSpace(row[1])}
		//result = append(result, quiz)
		// same as before we know the length no need to leave it to the heaps
		result[i] = quiz
	}

	return result
}

// PresentQuiz a function that allows you to run the quiz
func PresentQuiz(quiz *[]Quiz, timeout int) {
	input := bufio.NewScanner(os.Stdin)
	total := len(*quiz)
	fmt.Printf("Starting the quiz: %d questions\n\n", total)
	for i, q := range *quiz {
		fmt.Printf("Question %d of %d\n", i+1, total)
		fmt.Println(q.Question)
		fmt.Print("answer >  ")
		input.Scan()
		wasCorrect := (*quiz)[i].registerAnswer(input.Text())
		if wasCorrect {
			fmt.Print("Correct\n\n")
		} else {
			fmt.Printf("WRONG! It was %s\n\n", q.CorrectAnswer)
		}
	}
}

// ShowResult a function that calculates and shows the final result of the quiz
func ShowResult(quiz *[]Quiz) {
	correctAnswers := 0

	for _, q := range *quiz {
		if q.Answer == q.CorrectAnswer {
			correctAnswers++
		}
		println(q.ID, q.Question, q.Answer, q.CorrectAnswer)
	}

	total := len(*quiz)
	score := float32(correctAnswers) / float32(total) * 100.

	println()
	fmt.Printf("Questions %d | Correct %d   score: %d %%\n", total, correctAnswers, int32(score))
}
