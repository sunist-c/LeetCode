package day1

import (
	"testing"
	"time"
)

func judge(nums []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			return false
		}
	}

	return true
}

func TestSortedArrayDuplicates(t *testing.T) {
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
			args: args{nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
			want: 1,
		},
		{
			name: "2",
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			if got := SortedArrayDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("SortedArrayDuplicates() = %v, want %v", got, tt.want)
			}
			end := time.Now()
			if end.Sub(start) > time.Second {
				t.Errorf("time out")
			}
			if !judge(tt.args.nums, tt.want) {
				t.Errorf("array not duplicated completed")
			}
		})
	}
}
