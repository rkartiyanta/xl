package main

import (
	"fmt"

	"github.com/rkartiyanta/xl/answer"
)

func main() {
	question2 := []int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6}
	fmt.Printf("Inputs : %d\n", question2)
	fmt.Println("Result :")
	fmt.Println("number --> total")
	keys, results := answer.CheckNumberFrequency(question2)
	for _, v := range keys {
		fmt.Printf("%d -> %d\n", v, results[v])
	}
}
