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

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		// {"x", args{[]int{1, 1}, []int{1, 1}}, 1.5},
		// {"x", args{[]int{1, 1}, []int{1, 1}}, 1.5},
		{"x", args{[]int{1, 2}, []int{3, 4}}, 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"xx", args{"babad"}, "bab"},
		{"xx", args{"cbbd"}, "bb"},
		{"xx", args{"aba"}, "aba"},
		{"xx", args{"abba"}, "abba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"xx", args{"LEETCODEISHIRING", 3}, "LCIRETOESIIGEDHN"},
		{"xx", args{"LEETCODEISHIRING", 4}, "LDREOEIIECIHNTSG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"x", args{123}, 321},
		{"x", args{-123}, -321},
		{"x", args{120}, 21},
		{"x", args{123}, 321},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"xx", args{"   -42"}, -42},
		{"xx", args{"   +42"}, 42},
		{"xx", args{"42"}, 42},
		{"xx", args{"4193 with words"}, 4193},
		{"xx", args{"words and 987"}, 0},
		{"xx", args{"-91283472332"}, -2147483648},
		{"xx", args{"-000000000000000000000000000000000000000000000000001"}, -1},
		{"xx", args{"20000000000000000000"}, 2147483647},
		// """"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
