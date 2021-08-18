package answer

import (
	"testing"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name           string
		inputA         string
		inputB         string
		expectedResult string
		isSuccess      bool
	}{
		{
			name:           "Success - input A : saya, input B : kamu",
			inputA:         "saya",
			inputB:         "kamu",
			expectedResult: "skaaymau",
			isSuccess:      true,
		},
		{
			name:           "Success - input A : kosong, input B : ada",
			inputA:         "kosong",
			inputB:         "ada",
			expectedResult: "kaodsaong",
			isSuccess:      true,
		},
		{
			name:           "Success - input A : ada, input B : kosong",
			inputA:         "ada",
			inputB:         "kosong",
			expectedResult: "akdoasong",
			isSuccess:      true,
		},
		{
			name:           "Failed - input A : ada, input B : kosong",
			inputA:         "ada",
			inputB:         "kosong",
			expectedResult: "akdoasonggg",
			isSuccess:      false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := Concat(tc.inputA, tc.inputB)
			if tc.isSuccess == (res != tc.expectedResult) {
				t.Errorf("expected result is : %s, got %s",
					tc.expectedResult, res,
				)
			}
		})
	}
}
