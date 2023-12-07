package leetcode

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 1, 0, 1},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			},
			want: 5,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: 2,
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{1, 1, 0, 0, 1, 1, 1, 0, 1},
			},
			want: 4,
		},
        {
			name: "Example 5",
			args: args{
				nums: []int{1,0,1,0,1,0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
