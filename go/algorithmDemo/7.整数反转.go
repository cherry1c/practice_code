/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	var (
		result = 0
	)
	// x为正负数不用分类讨论。
	// 因为：-123 % 10 = -3，与正数的处理逻辑一致
	for x != 0 {
		mod := x % 10
		newResult := result*10 + mod
		// 倒推思路： C = A + B，如果没有溢出 A = C - B，溢出之后 A != C - B
		if (newResult-mod)/10 != result {
			return 0
		}

		result = newResult

		// 去掉最后一位
		x = x / 10

	}
	return result
}

// @lc code=end

