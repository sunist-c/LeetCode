package day2

import (
	"reflect"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test-9",
			args: args{
				grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				k:    1,
			},
			want: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name: "test-12",
			args: args{
				grid: [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
				k:    4,
			},
			want: [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShiftGrid(tt.args.grid, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShiftGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
