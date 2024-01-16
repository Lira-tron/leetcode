package leetcode

import (
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Example 1",
			args: args{
				path: "input.txt",
			},
		},
		{
			name: "Example 2",
			args: args{
				path: "input2.txt",
			},
		},
		{
			name: "Example 3",
			args: args{
				path: "input3.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			solve(tt.args.path)
		})
	}
}
