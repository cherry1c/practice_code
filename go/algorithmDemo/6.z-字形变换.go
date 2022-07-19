/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	rowsStr := make([][]rune, numRows)
	direction := 1
	index := 0
	for _, val := range s {
		rowsStr[index] = append(rowsStr[index], val)
		index += direction
		if index == numRows-1 {
			direction = -1
		} else if index == 0 {
			direction = 1
		}
	}
	var result string
	for _, val := range rowsStr {
		result += string(val)
	}
	return result
}

// @lc code=end

