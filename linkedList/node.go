package linkedList

import (
	"fmt"
)

type nodeRotationErrorType string

type nodeRotationErrorData struct {
	value nodeRotationErrorType
	msg   string
}

type nodeRotationError *nodeRotationErrorData

const (
	mainNodeNil       nodeRotationErrorType = "Main node is nil!"
	nodeCountExceeded nodeRotationErrorType = "Node count exceeded!"
)

// No support for generics, it is still in draft phase :(
// Draft as per v1.15: https://blog.golang.org/why-generics

type node struct {
	data int
	next *node
	prev *node
}

func (n *node) findMiddle() *node {
	if n == nil {
		return nil
	}

	slow := n
	fast := n

	for fast.next != nil {
		slow = slow.next
		fast = fast.next

		if fast.next != nil {
			fast = fast.next
		}
	}

	return slow
}

func (n *node) rotateCounterClockwise(count int) (*node, nodeRotationError) {
	if n == nil {
		return nil, &nodeRotationErrorData{
			value: mainNodeNil, msg: "",
		}
	}
	if count == 0 {
		return n, nil
	}

	var prev *node
	newHead := n
	index := 0
	requiredNodeReached := false
	for newHead.next != nil {
		if index == count-1 {
			prev = newHead
		}
		if index == count {
			requiredNodeReached = true
			break
		}
		newHead = newHead.next
		index++
	}

	availableNodeCount := index + 1
	if !requiredNodeReached && availableNodeCount < count {
		return nil, &nodeRotationErrorData{
			value: nodeCountExceeded,
			msg: fmt.Sprintf(
				"Rotation count %d > available Node count %d", count, availableNodeCount),
		}
	}

	// Disconnect node before newHead from pointing next to newHead
	// because it will be newTail
	if prev != nil {
		prev.next = nil
	}

	oldTail := newHead
	for oldTail.next != nil {
		oldTail = oldTail.next
	}

	// Attach the oldHead as next of the oldTail
	oldTail.next = n

	return newHead, nil
}

func (n *node) traverseToEnd() []int {
	values := make([]int, 0)

	curr := n
	for curr != nil {
		values = append(values, curr.data)
		curr = curr.next
	}

	return values
}

func (n *node) reverseInGroups(groupSize uint) *node {
	if groupSize == 0 {
		return n
	}

	var prev *node = nil
	curr := n
	next := n.next

	var index uint = 1
	for next != nil {
		if index >= groupSize {
			break
		}

		curr.next = prev

		prev = curr
		curr = next
		next = next.next

		index++
	}

	curr.next = prev

	if next != nil {
		n.next = next.reverseInGroups(groupSize)
	}

	return curr
}
