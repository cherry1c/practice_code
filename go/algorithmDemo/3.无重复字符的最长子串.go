/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	result := 0
	charMap := make(map[uint8]int)
	for left, right := 0, 0; right < len(s); right++ {
		c := s[right]
		val, ok := charMap[c]
		if !ok {
			val = 0
		}
		left = Max(val, left)
		result = Max(result, right-left+1)
		charMap[c] = right + 1
	}
	return result
}

func Max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

// @lc code=end

