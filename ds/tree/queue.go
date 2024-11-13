package main

type Queue []*Node

func (q *Queue) push(node *Node) {
	if node == nil {
		return
	}
	*q = append(*q, node)
}

func (q *Queue) pop() {
	if q.empty() {
		return
	}
	*q = (*q)[1:]
}

func (q *Queue) empty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

func (q *Queue) front() *Node {
	if len(*q) == 0 {
		return nil
	}

	return (*q)[0]
}
