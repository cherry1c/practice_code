package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		m[val] = i
	}
	for i, val := range nums {
		_, ok := m[target-val]
		if ok && i != m[target-val] {
			return []int{i, m[target-val]}
		}
	}
	return nil
}

// @lc code=end
