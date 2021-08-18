package main

import (
	"fmt"

	"github.com/rkartiyanta/xl/answer"
)

func main() {
	fmt.Println("String 1 : saya")
	fmt.Println("String 2 : kamu")
	fmt.Printf("Result : %s\n", answer.Concat("saya", "kamu"))
	fmt.Println("String 1 : kosong")
	fmt.Println("String 2 : ada")
	fmt.Printf("Result : %s\n", answer.Concat("kosong", "ada"))
	fmt.Println("String 1 : ada")
	fmt.Println("String 2 : kosong")
	fmt.Printf("Result : %s\n", answer.Concat("ada", "kosong"))
}
