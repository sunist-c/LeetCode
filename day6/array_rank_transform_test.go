package day6

import (
	"reflect"
	"testing"
)

func TestArrayRankTransform(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_case",
			args: args{arr: []int{40, 10, 20, 30}},
			want: []int{4, 1, 2, 3},
		},
		{
			name: "test_case_2",
			args: args{arr: []int{37, 12, 28, 9, 100, 56, 80, 5, 12}},
			want: []int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
		{
			name: "test_case_3",
			args: args{arr: []int{100, 100, 100}},
			want: []int{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayRankTransform(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
