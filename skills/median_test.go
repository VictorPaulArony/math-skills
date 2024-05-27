package skills

import "testing"

func TestMedian(t *testing.T) {
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
			args: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}
