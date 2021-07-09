/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	cnt := 0
	for i := 0; i < len(flowerbed); {
		if flowerbed[i] == 0 {
			if i+1 < len(flowerbed) && flowerbed[i+1] == 0 && (i-1 < 0 || flowerbed[i-1] == 0) ||
				i+1 == len(flowerbed) && (i-1 < 0 || flowerbed[i-1] == 0) {
				cnt++
				i += 2
				continue
			}
		}
		i++
	}

	if cnt >= n {
		return true
	}
	return false
}

// @lc code=end

