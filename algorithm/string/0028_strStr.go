package string

import "strings"

// 解法一
func StrStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

// 解法二
func StrStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
