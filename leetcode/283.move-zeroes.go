/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
	pos := []int{}
	for i:=0; i<len(nums); i++ {
		if nums[i] != 0 {
			pos = append(pos, i)
		}
	}
	for i,j := range pos{
		nums[i] = nums[j]
	}
	for i:=len(pos); i<len(nums); i++ {
		nums[i] = 0
	}
}
// @lc code=end

