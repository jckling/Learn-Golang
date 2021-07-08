/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

// @lc code=start
import "sort"

func thirdMax(nums []int) int {
	n := len(nums) - 1
	sort.Ints(nums)

	a := nums[n]
	for i := n - 1; i >= 0; i-- {
		b := nums[i]
		if b == a {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			c := nums[j]
			if c == b {
				continue
			}
			return c
		}
	}

	return a
}

// @lc code=end

