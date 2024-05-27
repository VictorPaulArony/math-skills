package skills

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name string
		args []float64
		want float64
	}{
		{
			name: "test1",
			args: []float64{2, 3, 4, 5},
			want: 1.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}
