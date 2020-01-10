package main

import (
	"strings"
)

var input = [...]string{"string", "to", "join", "to", "something", "like", "this", "text", "here!"}

// Replay command line argumants using ' ' as a seperator.
func main() {
	EchoRange()
	EchoJoin()
}

func EchoRange() string {
	s, sep := "", ""
	//	fmt.Println("Command Line: " + os.Args[0])

	for _, arg := range input[1:] {
		s += sep + arg
		sep = " "
	}

	return s
	//fmt.Println(s)
}

func EchoJoin() string {
	s, sep := "", " "
	//	fmt.Println("Command Line: " + os.Args[0])

	s = strings.Join(input[1:], sep)
	return s
	//fmt.Println(s)
}
