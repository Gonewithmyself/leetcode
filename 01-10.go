package leetcode

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
		m         = make(map[byte]struct{}, 1024)
	)

	for i < l && j < l {
		if _, ok := m[s[j]]; !ok {
			m[s[j]] = struct{}{}
			j++
			if cnt < j-i {
				cnt = j - i
			}
		} else {
			delete(m, s[i])
			i++
		}
	}

	return cnt
}
