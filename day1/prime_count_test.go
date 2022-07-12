package day1

import (
	"testing"
	"time"
)

func TestPrimeCount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100",
			args: args{n: 100},
			want: 25,
		},
		{
			name: "2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "1000",
			args: args{n: 1000},
			want: 168,
		},
		{
			name: "10000000",
			args: args{n: 10000000},
			want: 664579,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			if got := PrimeCount(tt.args.n); got != tt.want {
				t.Errorf("PrimeCount() = %v, want %v", got, tt.want)
			}
			end := time.Now()
			if end.Sub(start) > time.Second {
				t.Errorf("time out")
			}
		})
	}
}
