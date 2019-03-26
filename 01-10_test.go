package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"x", args{genListNode(2, 4, 3), genListNode(5, 6, 4)}, genListNode(7, 0, 8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				//printList(got)
				printList(tt.want)
				printList(got)
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Error("xx")

}

func genListNode(args ...int) *ListNode {
	h := &ListNode{}
	cur := h
	for i := range args {
		cur.Next = &ListNode{
			Val: args[i],
		}

		cur = cur.Next
	}

	return h.Next
}

func printList(l *ListNode) {
	for nil != l {
		fmt.Print(l.Val, "->")
		l = l.Next
	}

	fmt.Print("nil\n")
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{"xx", " ", 1},
		{"xx", "", 0},
		{"xu", "xu", 2},
		{"cdd", "cdd", 2},
		{"xx", "xx", 1},
		{"xx", "xyxx", 2},
		{"xz", "xyxzx", 3},

		{"xx", "abcabcbb", 3},
		{"pwwkew", "pwwkew", 3},
		// {"xx", ss, 95},
		//{"abcabcbb"}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
