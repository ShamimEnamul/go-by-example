package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	numOfNodes := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		numOfNodes++
	}

	if n == numOfNodes {
		return head.Next
	}

	temp = head
	for i := 1; temp != nil && i < numOfNodes-1; i++ {
		temp = temp.Next
	}

	temp = temp.Next.Next

	return head

}

func main() {

}
