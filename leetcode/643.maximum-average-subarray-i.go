/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	max := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		if i < k {
			sum += nums[i]
			if i == k-1 {
				max = sum
			}
		} else {
			sum = sum - nums[i-k] + nums[i]
		}

		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

// @lc code=end

