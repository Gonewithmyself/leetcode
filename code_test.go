package main

import "testing"

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
