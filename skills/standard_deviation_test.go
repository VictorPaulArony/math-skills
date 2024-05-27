package skills

import (
	"math"
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{
			name:     "positive numbers",
			input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 2.581988897471611, // Calculated standard deviation
		},
		{
			name:     "single number",
			input:    []float64{5},
			expected: 0,
		},
		{
			name:     "empty slice",
			input:    []float64{},
			expected: 0,
		},
		{
			name:     "same numbers",
			input:    []float64{5, 5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "negative numbers",
			input:    []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9},
			expected: 2.581988897471611, // Calculated standard deviation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StandardDeviation(tt.input)
			if math.Abs(got-tt.expected) > 1e-9 { // Allowing a small margin of error for floating-point comparisons
				t.Errorf("StandardDeviation() = %v, want %v", got, tt.expected)
			}
		})
	}
}
