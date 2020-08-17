package main

import "testing"

func Test_fizzbuzz(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "returnValue",
			args: args{
				x: 1,
			},
			want: "1",
		},
		{
			name: "returnFizz",
			args: args{
				x: 3,
			},
			want: "Fizz",
		},
		{
			name: "returnBuzz",
			args: args{
				x: 10,
			},
			want: "Buzz",
		},
		{
			name: "returnFizzBuzz",
			args: args{
				x: 30,
			},
			want: "FizzBuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzz(tt.args.x); got != tt.want {
				t.Errorf("fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
