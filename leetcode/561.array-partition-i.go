/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] Array Partition I
 */

// @lc code=start
import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

// @lc code=end

