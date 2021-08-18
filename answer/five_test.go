package answer

import (
	"testing"
)

func TestFindMaxValue(t *testing.T) {
	tests := []struct {
		name           string
		inputs         []string
		expectedResult int
		isSuccess      bool
	}{
		{
			name: "Success - Test A",
			inputs: []string{
				"5 3",
				"1 2 100",
				"2 5 100",
				"3 4 100",
			},
			expectedResult: 200,
			isSuccess:      true,
		},
		{
			name: "Success - Test B",
			inputs: []string{
				"5 3",
				"1 2 100",
				"2 5 100",
				"2 4 100",
			},
			expectedResult: 300,
			isSuccess:      true,
		},
		{
			name: "Failed - Test C",
			inputs: []string{
				"0 3",
				"1 2 100",
				"2 5 100",
				"2 4 100",
			},
			expectedResult: 300,
			isSuccess:      false,
		},
		{
			name: "Failed - Test D",
			inputs: []string{
				"5 0",
				"1 2 100",
				"2 5 100",
				"2 4 100",
			},
			expectedResult: 300,
			isSuccess:      false,
		},
		{
			name: "Failed - Test E",
			inputs: []string{
				"5 3",
				"1 2 100",
				"2 5 100",
				"6 4 100",
			},
			expectedResult: 300,
			isSuccess:      false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := FindMaxValue(tc.inputs)
			if tc.isSuccess == (tc.expectedResult != result) {
				t.Errorf("Expected result is %d, got %d", tc.expectedResult, result)
			}
		})
	}
}
