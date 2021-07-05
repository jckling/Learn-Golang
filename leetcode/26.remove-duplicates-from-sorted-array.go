/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r := 0, 1
	for ; r < len(nums); r++ {
		if nums[r] == nums[l] {
			continue
		}
		nums[l+1] = nums[r]
		l++
	}
	return l + 1
}

// @lc code=end

