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
	for next != nil && index < groupSize {
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

func (n *node) computeLength() uint {
	var size uint = 0
	curr := n

	for curr != nil {
		size++
		curr = curr.next
	}

	return size
}

func (n *node) traverse(count uint) *node {
	if count == 0 {
		return n
	}

	var index uint = 0
	curr := n

	for curr.next != nil && index < count {
		curr = curr.next
		index++
	}

	return curr
}

func (n *node) findIntersection(headB *node) *node {
	if headB == nil {
		return nil
	}

	sizeA := n.computeLength()
	sizeB := headB.computeLength()
	var minSize uint
	if sizeA < sizeB {
		minSize = sizeA
	} else {
		minSize = sizeB
	}

	currA := n.traverse(sizeA - minSize)
	currB := headB.traverse(sizeB - minSize)
	for currA != nil && currB != nil {
		if currA.next == currB.next {
			return currA.next
		}
		currA = currA.next
		currB = currB.next
	}

	return nil
}

func (n *node) isLoopPresent() bool {
	slow := n
	fast := n

	for fast != nil {
		slow = slow.next
		fast = fast.next

		if fast != nil {
			fast = fast.next
		}

		if slow == fast {
			return true
		}
	}

	return false
}

func (n *node) removeLoop() {
	slow := n
	fast := n

	for fast != nil {
		slow = slow.next
		fast = fast.next

		if fast != nil {
			fast = fast.next
		}

		if slow == fast {
			break
		}
	}

	if slow == nil || fast == nil {
		return
	}

	var prev *node
	curr := n
	for curr != slow {
		prev = slow
		curr = curr.next
		slow = slow.next
	}

	if prev != nil {
		prev.next = nil
	}
}

func (n *node) getNthFromLast(index uint) *int {
	first := n
	curr := n

	var count uint = 0
	for ; count < index; count++ {
		if curr.next == nil {
			return nil
		}

		curr = curr.next
	}

	for curr.next != nil {
		first = first.next
		curr = curr.next
	}

	return &first.data
}
