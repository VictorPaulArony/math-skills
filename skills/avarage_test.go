package skills

import "testing"

func TestAvarage(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name     string
		data     []float64
		expected float64
	}{
		{
			name:     "Test1",
			data:     []float64{2, 3},
			expected: 2.5,
		},

		{
			name:     "Test2",
			data:     []float64{10, 20, 50, 70},
			expected: 37.5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Avarage(test.data)
			if got != test.expected {
				t.Errorf("Avarage() = %v, want %v", got, test.expected)
			}
		})
	}
}
