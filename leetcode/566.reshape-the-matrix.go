/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] Reshape the Matrix
 */

// @lc code=start
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) == r*c {
		res := [][]int{}
		for i := 0; i < r; i++ {
			tmp := []int{}
			for j := 0; j < c; j++ {
				cnt := i*c + j
				tmp = append(tmp, mat[cnt/len(mat[0])][cnt%len(mat[0])])
			}
			res = append(res, tmp)
		}
		return res
	}
	return mat
}

// @lc code=end

