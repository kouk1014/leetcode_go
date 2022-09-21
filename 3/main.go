package main

import "strings"

func main() {
	lengthOfLongestSubstring("abcabcbb")
}
func lengthOfLongestSubstring(s string) int {
	subStr := ""
	max := 0
	for _, c := range s {
		s := string(c)
		if i := strings.Index(subStr, s); i >= 0 {
			subStr = subStr[i+1:]
		}
		subStr += s
		if len(subStr) > max {
			max = len(subStr)
		}
	}
	return max
}
