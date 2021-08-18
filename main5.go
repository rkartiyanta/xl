package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rkartiyanta/xl/answer"
)

const (
	minN = 3
	maxN = 10000000

	minM = 1
	maxM = 200000

	minAB = 1
	maxK  = 1000000000
)

func main() {
	question5 := make([]string, 0)
	// question5 := []string{
	// 	"5 3",
	// 	"1 2 100",
	// 	"2 5 100",
	// 	"3 4 100",
	// }
	// fmt.Printf("Result : %d\n", answer.FindmaxK(question5))

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Number Five")
	fmt.Println("---------------------")
	i := 0
	n := 0
	m := 0
	var err error
	max := 99
	for i < max {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if i == 0 {
			params := strings.Split(input, " ")
			n, err = strconv.Atoi(params[0])
			if err != nil || n < minN || n > maxN {
				fmt.Printf("N must be between %d and %d\n", minN, maxN)
				return
			}

			m, err = strconv.Atoi(params[1])
			if err != nil || m < minM || m > maxM {
				fmt.Printf("M must be between %d and %d\n", minM, maxM)
				return
			}
			max = m + 1
		} else {
			input = strings.TrimSpace(input)
			if input == "" {
				fmt.Println("Test input cannot be empty")
				return
			}

			question := strings.Split(input, " ")
			a, err := strconv.Atoi(question[0])
			if err != nil || a < minAB {
				fmt.Printf("a(%d) must be bigger than %d\n", a, minAB)
				return
			}
			if a > n {
				fmt.Printf("a(%d) must be smaller than N(%d)\n", a, n)
				return
			}

			b, err := strconv.Atoi(question[1])
			if err != nil || b < minAB {
				fmt.Printf("b(%d) must be bigger than %d\n", b, minAB)
				return
			}
			if b > n {
				fmt.Printf("b(%d) must be smaller than N(%d)\n", b, n)
				return
			}

			if b < a {
				fmt.Printf("b(%d) must be bigger than a(%d)\n", b, a)
				return
			}

			k, err := strconv.Atoi(question[2])
			if err != nil {
				fmt.Println("k is not a number")
				return
			}
			if k > maxK {
				fmt.Printf("k(%d) must be smaller than N(%d)\n", k, maxK)
				return
			}
		}
		question5 = append(question5, input)
		i++
	}
	fmt.Println("---------------------")
	fmt.Printf("Result : %d\n", answer.FindMaxValue(question5))
}
