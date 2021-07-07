/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// @lc code=end

