package day1

import (
	"testing"
	"time"
)

func TestPivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{nums: []int{1, 7, 3, 6, 5, 6}},
			want: 3,
		},
		{
			name: "2",
			args: args{nums: []int{1, 2, 3}},
			want: -1,
		},
		{
			name: "3",
			args: args{nums: []int{2, 1, -1}},
			want: 0,
		},
		{
			name: "4",
			args: args{nums: []int{-1, -1, 0, 1, 1, 0}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			if got := PivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("PivotIndex() = %v, want %v", got, tt.want)
			}
			end := time.Now()
			if end.Sub(start) > time.Second {
				t.Errorf("time out")
			}
		})
	}
}
