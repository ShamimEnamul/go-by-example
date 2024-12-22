package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	len := 0
	for curr := head; curr != nil; curr = curr.Next {
		len++
	}
	fmt.Println(len)
	if k == 0 || len == 0 {
		return head
	}
	k = k % len
	if k == 0 {
		return head
	}
	step := len - k
	curr := head
	for i := 1; i < step; i++ {
		curr = curr.Next
	}
	next := curr.Next
	curr.Next = nil
	for curr = next; curr.Next != nil; curr = curr.Next {
	}
	curr.Next = head
	return next
}
func main() {
	fmt.Println("Rotate a linked list")
}
