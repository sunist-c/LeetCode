package day1

import "testing"

func TestSqrtWithBinary(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{x: 4},
			want: 2,
		},
		{
			name: "2",
			args: args{x: 8},
			want: 2,
		},
		{
			name: "3",
			args: args{x: 1919810},
			want: 1385,
		},
		{
			name: "4",
			args: args{x: 114514},
			want: 338,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SqrtWithBinary(tt.args.x); got != tt.want {
				t.Errorf("SqrtWithBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqrtWithNewton(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{x: 4},
			want: 2,
		},
		{
			name: "2",
			args: args{x: 8},
			want: 2,
		},
		{
			name: "3",
			args: args{x: 1919810},
			want: 1385,
		},
		{
			name: "4",
			args: args{x: 114514},
			want: 338,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SqrtWithNewton(tt.args.x); got != tt.want {
				t.Errorf("SqrtWithNewton() = %v, want %v", got, tt.want)
			}
		})
	}
}
