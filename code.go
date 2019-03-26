package main

func main() {
	lengthOfLongestSubstring("PWD")
}

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
