/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] == val {
			continue
		}
		nums[l] = nums[r]
		l++
	}
	return l
}

// @lc code=end

