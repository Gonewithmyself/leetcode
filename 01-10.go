package leetcode

import (
	"bytes"
)

// 001
func twoSum(nums []int, target int) []int {
	l := len(nums)
	m := make(map[int]int, l)
	for i := range nums {
		c := target - nums[i]
		if idx, ok := m[c]; ok {
			return []int{idx, i}
		}
		m[nums[i]] = i
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 002
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		h      = &ListNode{}
		cur    = h
		f, sum int
	)
	for nil != l1 || nil != l2 {
		if nil != l1 {
			sum += l1.Val
			l1 = l1.Next
		}

		if nil != l2 {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += f
		if sum > 9 {
			sum -= 10
			f = 1
		} else {
			f = 0
		}

		cur.Next = &ListNode{
			Val: sum,
		}
		cur = cur.Next
		sum = 0
	}

	if 0 != f {
		cur.Next = &ListNode{
			Val: f,
		}

	}
	return h.Next
}

// 003
func lengthOfLongestSubstring(s string) int {

	var (
		i, j, cnt int
		l         = len(s)
		m         = make(map[byte]struct{}, 512)
	)

	for i < l && j < l {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = struct{}{}
			i++
			if cnt < i-j {
				cnt = i - j
			}
		} else {
			delete(m, s[j])
			j++
		}
	}
	return cnt

}

// 004
func findMedianSortedArrays(A []int, B []int) float64 {
	var (
		m = len(A)
		n = len(B)
	)

	if m > n {
		A, B, m, n = B, A, n, m
	}

	var (
		imin, imax, half = 0, m, (m + n + 1) / 2
		i, j             int
	)

	for imin <= imax {
		i = (imin + imax) / 2
		j = half - i
		if i < m && B[j-1] > A[i] {
			imin = i + 1
		} else if i > 0 && A[i-1] > B[j] {
			imax = i - 1
		} else {
			var left, right int
			if i == 0 {
				left = B[j-1]
			} else if j == 0 {
				left = A[i-1]
			} else {
				left = A[i-1]
				if left < B[j-1] {
					left = B[j-1]
				}
			}

			if (m+n)&1 != 0 {
				return float64(left)
			}

			if i == m {
				right = B[j]
			} else if j == n {
				right = A[i]
			} else {
				right = A[i]
				if right > B[j] {
					right = B[j]
				}
			}

			return (float64(right) + float64(left)) / 2
		}
	}

	return 0
}

// 005
func longestPalindrome(s string) string {
	var res, tmp string
	for i := range s {
		tmp = longestPalindromeHelper(s, i, i)
		if len(tmp) > len(res) {
			res = tmp
		}
		tmp = longestPalindromeHelper(s, i, i+1)
		if len(tmp) > len(res) {
			res = tmp
		}
	}

	return res
}

func longestPalindromeHelper(s string, l, r int) string {
	c := len(s)
	for l >= 0 && r < c && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

// 006
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var (
		rows = make([][]byte, numRows)
		idx  int
		down = true
	)
	for i := 0; i < numRows; i++ {
		rows[i] = make([]byte, 0, 1024)
	}

	for i := range s {
		rows[idx] = append(rows[idx], s[i])
		if down {
			idx++
		} else {
			idx--
		}
		if idx > numRows-1 {
			idx -= 2
			down = !down
		} else if idx < 0 {
			idx += 2
			down = !down
		}
	}

	return string(bytes.Join(rows, nil))
}
