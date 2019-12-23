package action

import "testing"

func Test_averageOfTwo(t *testing.T) {
	type args struct {
		a float32
		b float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"1", args{1, 1}, 1},
		{"2", args{1.5, 1}, 1.25},
		{"3", args{1002, 1}, 501.5},
		{"4", args{3, 4}, 3.5},
		{"5", args{0, 0}, 0},
		{"6", args{0, 1}, .5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfTwo(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("averageOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
