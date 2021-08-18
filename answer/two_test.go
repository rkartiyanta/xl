package answer

import (
	"testing"
)

func TestCheckNumberFrequency(t *testing.T) {
	tests := []struct {
		name          string
		inputs        []int
		expectedKey   []int
		expectedValue []int
	}{
		{
			name:          "Input : [4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6]",
			inputs:        []int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6},
			expectedKey:   []int{4, 6, 5, 3, 7, 8},
			expectedValue: []int{6, 5, 3, 2, 2, 1},
		},
		{
			name:          "Input : [4, 6, 3, 5, 4, 6, 7, 7, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6]",
			inputs:        []int{4, 6, 3, 5, 4, 6, 7, 7, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6},
			expectedKey:   []int{4, 6, 5, 7, 3},
			expectedValue: []int{6, 5, 3, 3, 2},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keys, results := CheckNumberFrequency(tc.inputs)
			for i := range keys {
				if keys[i] != tc.expectedKey[i] {
					t.Errorf("Expected key %d, got %d",
						tc.expectedKey, keys,
					)
				}
				if results[keys[i]] != tc.expectedValue[i] {
					t.Errorf("Expected value for %d : %d, got %d",
						keys[i], tc.expectedValue[i], results[keys[i]],
					)
				}
			}
		})
	}
}
