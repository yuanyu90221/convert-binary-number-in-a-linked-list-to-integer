package decimal_linkedlist

import (
	"testing"
)

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				head: BuildList([]int{1, 0, 1}),
			},
			want: 5,
		},
		{
			name: "Example2",
			args: args{
				head: BuildList([]int{0}),
			},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				head: BuildList([]int{1}),
			},
			want: 1,
		},
		{
			name: "Example4",
			args: args{
				head: BuildList([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}),
			},
			want: 18880,
		},
		{
			name: "Example5",
			args: args{
				head: BuildList([]int{0, 0}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BuildList(nums []int) *ListNode {
	var result *ListNode
	var currentNode *ListNode
	for idx, val := range nums {
		tempNode := ListNode{Val: val, Next: nil}
		if idx == 0 {
			result = &tempNode
			currentNode = &tempNode
		} else {
			currentNode.Next = &tempNode
			currentNode = currentNode.Next
		}
	}
	return result
}
