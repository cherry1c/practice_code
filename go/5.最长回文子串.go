/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	sLen := len(s)
	if slen == 0 {
		return s
	}
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}
	maxLen := 0
	result := ""

	for i := 0; i < sLen; i++ {
		for j := 0; j < i; j++ {
			if s[j] != s[i] {
				dp[i][j] = false
				continue
			}
			if (i-j <= 1) || (i > 0 && s[i-1] == s[j+1]) {
				dp[i][j] = true
				if maxLen < i-j {
					result = s[j : i+1]
				}
				continue
			} else {
				dp[i][j] = false
			}
		}
	}
}

// @lc code=end

