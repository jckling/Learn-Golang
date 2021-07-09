/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] Maximum Product of Three Numbers
 */

// @lc code=start
import "sort"

func maximumProduct(nums []int) int {
	n := len(nums)
	product := 1
	if n <= 3 {
		for _, x := range nums {
			product *= x
		}
	} else {
		sort.Ints(nums)
		l, r := nums[0]*nums[1], nums[n-1]*nums[n-2]
		a, b, c := l*nums[2], l*nums[n-1], r*nums[n-3]
		if a > b && a > c {
			return a
		} else if b > a && b > c {
			return b
		} else {
			return c
		}
	}

	return product
}

// @lc code=end

