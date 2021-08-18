package answer

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	minAnswer = 3
	maxAnswer = 10000000

	minQuestion = 1
	maxQuestion = 200000

	minAB    = 1
	maxValue = 1000000000
)

func FindMaxValue(inputs []string) int {
	max := 0
	firstParams := strings.Split(inputs[0], " ")
	lenAnswer, err := strconv.Atoi(firstParams[0])
	if err != nil || lenAnswer < minAnswer || lenAnswer > maxAnswer {
		fmt.Printf("N must be between %d and %d\n", minAnswer, maxAnswer)
		return max
	}

	lenQuestion, err := strconv.Atoi(firstParams[1])
	if err != nil || lenQuestion < minQuestion || lenQuestion > maxQuestion {
		fmt.Printf("M must be between %d and %d\n", minQuestion, maxQuestion)
		return max
	}
	if lenQuestion != len(inputs)-1 {
		fmt.Println("Question missmatch")
		return max
	}
	answers := make([]int, lenAnswer)

	for i := 1; i < len(inputs); i++ {
		question := strings.Split(inputs[i], " ")
		firstIndex, err := strconv.Atoi(question[0])
		if err != nil || firstIndex < minAB {
			fmt.Printf("a(%d) must be bigger than %d\n", firstIndex, minAB)
			return max
		}
		if firstIndex > lenAnswer {
			fmt.Printf("a(%d) must be smaller than N(%d)\n", firstIndex, lenAnswer)
			return max
		}

		secondIndex, err := strconv.Atoi(question[1])
		if err != nil || secondIndex < minAB {
			fmt.Printf("a(%d) must be bigger than %d\n", secondIndex, minAB)
			return max
		}
		if secondIndex > lenAnswer {
			fmt.Printf("a(%d) must be smaller than N(%d)\n", secondIndex, lenAnswer)
			return max
		}
		if secondIndex < firstIndex {
			fmt.Printf("b(%d) must be bigger than a(%d)\n", secondIndex, firstIndex)
			return max
		}
		value, err := strconv.Atoi(question[2])
		if err != nil {
			fmt.Println("value is not a number")
			return max
		}
		if value > maxValue {
			fmt.Printf("value(%d) must be smaller than N(%d)\n", value, maxValue)
			return max
		}
		for index := firstIndex - 1; index < secondIndex; index++ {
			answers[index] += value
		}
	}

	for i := range answers {
		if answers[i] > max {
			max = answers[i]
		}
	}

	return max
}
