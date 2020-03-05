package leetcode

import (
	"math"
)

// 020
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	match := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 64)
	pos := 0

	for i := range s {
		pop := false
		switch s[i] {
		case '(':
		case '[':
		case '{':
		default:
			pop = true
		}

		if pop {
			if pos == 0 {
				return false
			}

			pos--

			if match[stack[pos]] != byte(s[i]) {
				return false
			}

		} else {
			if pos >= len(stack) {
				extend := make([]byte, 64)
				stack = append(stack, extend...)
			}
			stack[pos] = byte(s[i])
			pos++
		}

	}

	if pos == 0 {
		return true
	}
	return false
}

// 026
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	j := 0
	for i := 0; i < n-1; i++ {
		k := i + 1
		if nums[k] != nums[i] {
			nums[j] = nums[i]
			j++
			continue
		}
	}

	nums[j] = nums[n-1]
	j++
	return j
}

// 53
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// maxsum := nums[0]
	// for i := range nums {
	// 	sum := 0
	// 	for j := i; j < n; j++ {
	// 		sum += nums[j]
	// 		maxsum = int(math.Max(float64(sum), float64(maxsum)))
	// 	}
	// }
	// return maxsum

	maxsum := nums[0]
	sum := 0
	for i := range nums {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		maxsum = int(math.Max(float64(sum), float64(maxsum)))
	}
	return maxsum
}

// 88
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	i := m - 1
	j := n - 1
	for ; k >= 0; k-- {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		} else if i >= 0 {
			nums1[k] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 121
func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	profit := 0
	for i := range prices {
		if curr := prices[i] - minPrice; curr > profit {
			profit = curr
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return profit
}

// 0125
func isPalindromeStr(s string) bool {
	j := len(s) - 1
	if j < 0 {
		return true
	}
	i := 0
	for j >= i {
		a, aok := validChar((s[i]))
		b, bok := validChar(s[j])
		if aok && bok {
			if a != b {
				return false
			}
			i++
			j--
		}

		if !aok {
			i++
		}

		if !bok {
			j--
		}
	}
	return true
}

func validChar(c byte) (byte, bool) {
	if c >= 'a' && c <= 'z' {
		c -= 'a' - 'A'
	}

	if c >= 'A' && c <= 'Z' || (c >= '0' && c <= '9') {
		return c, true
	}
	return c, false
}

// 0136
func singleNumber(nums []int) int {
	// m := make(map[int]struct{},8)
	// for i := range nums {
	// 	if _, ok := m[nums[i]]; ok {
	// 		delete(m, nums[i])
	// 	} else {
	// 		m[nums[i]] = struct{}{}
	// 	}
	// }

	// for k := range m {
	// 	return k
	// }
	// return 0

	res := 0
	for i := range nums {
		res ^= nums[i]
	}
	return res
}

func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}

		if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{-1, -1}
}

func rob(nums []int) int {
	// dp[i]= max(dp[i-2]+nums[i], dp[i-1])
	n := len(nums)
	dp := make([]int, n)
	switch n {
	case 0:
		return 0
	case 1:
		dp[0] = nums[0]
	default:
		dp[0] = nums[0]
		if nums[0] > nums[1] {
			dp[1] = nums[0]
		} else {
			dp[1] = nums[1]
		}
	}

	for i := 2; i < n; i++ {
		cur := dp[i-2] + nums[i]
		if cur >= dp[i-1] {
			dp[i] = cur
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}
