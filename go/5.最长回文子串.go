/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package main

// @lc code=start
func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return s
	}
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}

	result := s[0:1]

	for i := 0; i < sLen; i++ {
		dp[i][i] = true
		for j := 0; j < i; j++ {
			if s[j] == s[i] && (i-j <= 1 || dp[i-1][j+1]) {
				dp[i][j] = true
				if len(result) < i-j+1 {
					result = s[j : i+1]
				}
			}
		}
	}
	return result
}

// @lc code=end
