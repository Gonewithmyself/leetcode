package leetcode

import (
	"reflect"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{""}, true},
		{"2", args{"()"}, true},
		{"3", args{"(]"}, false},
		{"4", args{"()[]{}"}, true},
		{"5", args{"([)]"}, false},
		{"6", args{"{[]}"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, 5},
		{"2", args{[]int{1, 2}}, 2},
		{"3", args{[]int{1, 1, 2}}, 2},
		{"4", args{[]int{1, 2, 2}}, 2},
		{"4", args{[]int{1, 1, 2, 2, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"2", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name  string
		args  args
		wants []int
	}{
		// TODO: Add test cases.
		{"2", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}},
		{"3", args{[]int{2, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3}, []int{1, 2, 2, 3, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.wants) {
				t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.wants)
			}
		})
	}
}

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"2", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"1", args{[]int{7, 6, 4, 3, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_isPalindromeStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{"A man, a plan, a canal: Panama"}, true},
		{"2", args{"race a car"}, false},
		{"3", args{""}, true},
		{"3", args{"a"}, true},
		{"3", args{"ab"}, false},
		{"3", args{"ab a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeStr(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 0167
func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{2, 2, 1}}, 1},
		{"3", args{[]int{4, 1, 2, 1, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum2(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.

		{"1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{1, 2, 3, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
