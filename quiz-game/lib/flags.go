package lib

import "flag"

const defaultQuizFile = "problems.csv"

// GetArgs stuff thing
func GetArgs() (string, int) {
	quiz := flag.String("q", defaultQuizFile, "-q quiz file")
	timeout := flag.Int("t", -1, "-t timeout")
	flag.Parse()

	if *quiz == "" {
		panic("please specify quiz file -q filename")
	}

	return *quiz, *timeout
}
