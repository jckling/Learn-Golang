/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}
	for i, x := range nums {
		if x != i {
			return i
		}
	}
	return 0
}

// @lc code=end

