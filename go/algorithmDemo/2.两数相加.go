package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var carry int // 进位
	ans := ListNode{Val: 0, Next: nil}
	currentPtr := &ListNode{Val: 0, Next: &ans}
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		currentPtr = currentPtr.Next
		currentPtr.Val = (carry + p1.Val + p2.Val) % 10
		currentPtr.Next = &ListNode{Val: 0, Next: nil}
		carry = (carry + p1.Val + p2.Val) / 10
		p1 = p1.Next
		p2 = p2.Next
	}

	for p1 != nil {
		currentPtr = currentPtr.Next
		currentPtr.Val = (carry + p1.Val) % 10
		currentPtr.Next = &ListNode{Val: 0, Next: nil}
		carry = (carry + p1.Val) / 10
		p1 = p1.Next
	}

	for p2 != nil {
		currentPtr = currentPtr.Next
		currentPtr.Val = (carry + p2.Val) % 10
		currentPtr.Next = &ListNode{Val: 0, Next: nil}
		carry = (carry + p2.Val) / 10
		p2 = p2.Next
	}

	if carry != 0 {
		currentPtr = currentPtr.Next
		currentPtr.Val = carry
	}
	currentPtr.Next = nil
	return &ans
}

// @lc code=end
