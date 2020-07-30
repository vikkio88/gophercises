package main

import (
	"fmt"
	"quiz-game/lib"
)

func main() {
	quizfile, timeout := lib.GetArgs()
	quiz := lib.LoadQuiz(quizfile)
	println(fmt.Sprintf("timeout %d", timeout))
	lib.PresentQuiz(&quiz, timeout)

	lib.ShowResult(&quiz)
}
