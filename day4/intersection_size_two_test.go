package day4

import "testing"

func TestIntersectionSizeTwo(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}},
			want: 3,
		},
		{
			name: "test2",
			args: args{intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectionSizeTwo(tt.args.intervals); got != tt.want {
				t.Errorf("IntersectionSizeTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
