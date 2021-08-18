package answer

import (
	"fmt"
	"regexp"
)

var (
	regexAllowableInput = regexp.MustCompile(`^[A-Z]*$`)
	maxInput            = 100000
)

func NumberFour(inputs []interface{}) []string {
	result := make([]string, 0)
	tests := inputs[0].(int)
	if tests < 1 || tests > 5 {
		fmt.Println("Test case must between 1 and 5")
		return result
	}
	if len(inputs) != tests*2+1 {
		fmt.Println("Test input missmatch")
		return result
	}
	for i := 0; i < tests; i++ {
		firstWord := inputs[i*2+1].(string)
		if firstWord == "" {
			fmt.Println("Test input cannot be empty")
			return result
		}
		if firstWord == "" || !regexAllowableInput.MatchString(firstWord) || len(firstWord) > maxInput {
			fmt.Printf("Test input : %s, format error\n", firstWord)
			return result
		}
		secondWord := inputs[i*2+2].(string)
		if secondWord == "" {
			fmt.Println("Test input cannot be empty")
			return result
		}
		if secondWord == "" || !regexAllowableInput.MatchString(secondWord) || len(secondWord) > maxInput {
			fmt.Printf("Test input : %s, format error\n", secondWord)
			return result
		}
		result = append(result, concatLexicographically(firstWord, secondWord))
	}
	return result
}

func concatLexicographically(a, b string) string {
	var result string
	firstWords := getRunes(a)
	secondWords := getRunes(b)

	finish := false
	for !finish {
		if firstWords[0] > secondWords[0] {
			result = fmt.Sprintf("%s%c", result, secondWords[0])
			secondWords = removeFirst(secondWords)
		} else {
			result = fmt.Sprintf("%s%c", result, firstWords[0])
			firstWords = removeFirst(firstWords)
		}

		if len(firstWords) == 0 && len(secondWords) != 0 {
			result = concatRuneToString(result, secondWords)
			finish = true
		}
		if len(firstWords) != 0 && len(secondWords) == 0 {
			result = concatRuneToString(result, firstWords)
			finish = true
		}
	}

	return result
}

func concatRuneToString(result string, inputs []rune) string {
	for i := range inputs {
		result = fmt.Sprintf("%s%c", result, inputs[i])
	}

	return result
}

func removeFirst(s []rune) []rune {
	return append([]rune{}, s[1:]...)
}

func getRunes(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}
