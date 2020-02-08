package main

import (
	"fmt"
	"os"
)

// Fibonacci V1
// This aims to test and time different ways to generate the Fibonacci
// sequence.
// As of now, there is only one way that's been implemented and the
// code is still procedural as I have not learned go's OO model yet.
func main() {
	var limit int
	fmt.Println("FIBONACCI SEQUENCE")

	limit = readInt("How long should the series be?", limit)

	var seq []int
	for i := 0; i < limit; i++ {
		seq = append(seq, fibo(i))
	}
	fmt.Println(seq)
}

// The "original" recursive function, in a mathematical sense
// First few results are hard coded in.
func fibo(i int) int {
	switch {
	case i == 0:
		return 0
	case i < 3:
		return 1
	default:
		return fibo(i-1) + fibo(i-2)
	}
}

// Displays a prompt and reads and integer from the console
// Doesn't work very graciously yet
func readInt(msg string, i int) int {
	fmt.Println(msg)
	fmt.Print("(enter an integer between 1-49) :")
	_, err := fmt.Scanf("%d", &i)
	if err != nil || i > 49 {
		fmt.Println(err)
		os.Exit(1)
	}
	return i
}
