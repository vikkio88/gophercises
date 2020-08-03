package main

import (
	"quiz-game/lib"
)

func main() {
	quizfile, timeout := lib.GetArgs()
	quiz := lib.LoadQuiz(quizfile)
	lib.PresentQuiz(&quiz, timeout)

	lib.ShowResult(&quiz)
}
