/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
import "strconv"

func summaryRanges(nums []int) []string {
	lst := []string{}

	for i := 0; i < len(nums); i++ {
		l, r := i, i
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] == j-i {
				if j == len(nums)-1 {
					r = j
					i = j
					break
				}
				continue
			} else {
				r = j - 1
				i = j - 1
				break
			}
		}
		if l == r {
			lst = append(lst, strconv.Itoa(nums[l]))
		} else {
			lst = append(lst, strconv.Itoa(nums[l])+"->"+strconv.Itoa(nums[r]))
		}
	}
	return lst
}

// @lc code=end

