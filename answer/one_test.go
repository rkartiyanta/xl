package answer

import (
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		name      string
		inputA    int
		inputB    int
		expectedA int
		expectedB int
		isSuccess bool
	}{
		{
			name:      "Success - input A : 1, input B : 2",
			inputA:    1,
			inputB:    2,
			expectedA: 2,
			expectedB: 1,
			isSuccess: true,
		},
		{
			name:      "Failed - input A : 1, input B : 2",
			inputA:    1,
			inputB:    2,
			expectedA: 2,
			expectedB: 2,
			isSuccess: false,
		},
		{
			name:      "Success - input A : 10, input B : 20",
			inputA:    10,
			inputB:    20,
			expectedA: 20,
			expectedB: 10,
			isSuccess: true,
		},
		{
			name:      "Failed - input A : 10, input B : 20",
			inputA:    10,
			inputB:    20,
			expectedA: 20,
			expectedB: 20,
			isSuccess: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			Swap(&tc.inputA, &tc.inputB)
			if tc.isSuccess == (tc.inputA != tc.expectedA || tc.inputB != tc.expectedB) {
				t.Errorf("A should be %d and B should be %d, got A: %d, B: %d",
					tc.expectedA, tc.expectedB, tc.inputA, tc.inputB,
				)
			}
		})
	}
}
