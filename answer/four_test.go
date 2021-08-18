package answer

import (
	"testing"
)

func TestConcatLexicographically(t *testing.T) {
	tests := []struct {
		name           string
		inputA         string
		inputB         string
		expectedResult string
		isSuccess      bool
	}{
		{
			name:           "Success - input A : JACK, input B : DANIEL",
			inputA:         "JACK",
			inputB:         "DANIEL",
			expectedResult: "DAJACKNIEL",
			isSuccess:      true,
		},
		{
			name:           "Success - input A : ABACABA, input B : ABACABA",
			inputA:         "ABACABA",
			inputB:         "ABACABA",
			expectedResult: "AABABACABACABA",
			isSuccess:      true,
		},
		{
			name:           "Failed - input A : ABACABA, input B : ABACABA",
			inputA:         "ABACABA",
			inputB:         "ABACABA",
			expectedResult: "AABABACABACABAA",
			isSuccess:      false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := concatLexicographically(tc.inputA, tc.inputB)
			if tc.isSuccess == (tc.expectedResult != result) {
				t.Errorf("Expected result is %s, got %s", tc.expectedResult, result)
			}
		})
	}
}
