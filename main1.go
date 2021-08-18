package main

import (
	"fmt"

	"github.com/rkartiyanta/xl/answer"
)

func main() {
	a := 1
	b := 2
	fmt.Printf("Before a : %d, b : %d\n", a, b)
	answer.Swap(&a, &b)
	fmt.Printf("After a : %d, b : %d\n", a, b)
}
