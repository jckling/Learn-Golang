/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, m+n)
	i, j, k := 0, 0, 0
	for k < m+n {
		if i >= m {
			nums = append(nums[:k], nums2[j:]...)
			break
		}
		if j >= n {
			nums = append(nums[:k], nums1[i:]...)
			break
		}
		if nums1[i] <= nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}
	copy(nums1, nums)
}

// @lc code=end

