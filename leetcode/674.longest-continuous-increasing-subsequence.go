/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	max := 1
	cnt := 1
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if prev < nums[i] {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 1
		}
		prev = nums[i]
	}
	if cnt > max {
		max = cnt
	}
	return max
}

// @lc code=end

