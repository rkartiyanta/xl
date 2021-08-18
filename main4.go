package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rkartiyanta/xl/answer"
)

var (
	regexAllowableInput = regexp.MustCompile(`^[A-Z]*$`)
	maxInput            = 100000
)

func main() {
	question4 := make([]interface{}, 0)
	// question4 := []interface{}{
	// 	2,
	// 	"JACK",
	// 	"DANIEL",
	// 	"ABACABA",
	// 	"ABACABA",
	// }
	// answer.NumberFour(question4)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Number Four")
	fmt.Println("---------------------")
	i := 0
	max := 99
	for i < max {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if i == 0 {
			tempMax, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("First input must be a number")
				return
			}

			if tempMax < 1 || tempMax > 5 {
				fmt.Println("Test case must between 1 and 5")
				return
			}
			max = tempMax*2 + 1
			question4 = append(question4, tempMax)
		} else {
			input = strings.TrimSpace(input)
			if input == "" {
				fmt.Println("Test input cannot be empty")
				return
			}
			if input == "" || !regexAllowableInput.MatchString(input) || len(input) > maxInput {
				fmt.Printf("Test input : %s, format error\n", input)
				return
			}
			question4 = append(question4, input)
		}
		i++
	}
	fmt.Println("---------------------")
	fmt.Println("Result:")
	results := answer.NumberFour(question4)
	for _, v := range results {
		fmt.Println(v)
	}
}
