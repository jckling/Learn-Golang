/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] Image Smoother
 */

// @lc code=start
func imageSmoother(img [][]int) [][]int {
	smoothed := [][]int{}
	rows := len(img)
	cols := len(img[0])

	for r := 0; r < rows; r++ {
		tmp := []int{}
		for l := 0; l < cols; l++ {
			sum := img[r][l]
			cnt := 1
			if l+1 < cols {
				sum += img[r][l+1]
				cnt++
				if r-1 >= 0 {
					sum += img[r-1][l+1]
					cnt++
				}
				if r+1 < rows {
					sum += img[r+1][l+1]
					cnt++
				}
			}
			if l-1 >= 0 {
				sum += img[r][l-1]
				cnt++
				if r-1 >= 0 {
					sum += img[r-1][l-1]
					cnt++
				}
				if r+1 < rows {
					sum += img[r+1][l-1]
					cnt++
				}
			}
			if r-1 >= 0 {
				sum += img[r-1][l]
				cnt++
			}
			if r+1 < rows {
				sum += img[r+1][l]
				cnt++
			}
			tmp = append(tmp, sum/cnt)
		}
		smoothed = append(smoothed, tmp)
	}

	return smoothed
}

// @lc code=end

