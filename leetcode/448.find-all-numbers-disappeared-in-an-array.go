/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	res := []int{}

	for _, x := range nums {
		x = (x - 1) % len(nums)
		nums[x] = nums[x] + len(nums)
	}

	for i, x := range nums {
		if x <= len(nums) {
			res = append(res, i+1)
		}
	}

	return res
}

// @lc code=end

