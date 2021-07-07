/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	cnt := make(map[int]int)
	for _, x := range nums {
		cnt[x]++
	}
	for i := range cnt {
		if cnt[i] > len(nums)/2 {
			return i
		}
	}
	return 0
}

// @lc code=end

