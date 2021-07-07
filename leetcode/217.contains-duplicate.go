/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	cnt := make(map[int]int)
	for _, x := range nums {
		cnt[x]++
		if cnt[x] == 2 {
			return true
		}
	}
	for i := range cnt {
		if cnt[i] >= 2 {
			return true
		}
	}
	return false
}

// @lc code=end

