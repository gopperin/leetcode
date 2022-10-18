package array

import "github.com/gopperin/leetcode/algorithm/core"

func CanJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = core.Max(maxJump, i+v)
	}
	return true
}
