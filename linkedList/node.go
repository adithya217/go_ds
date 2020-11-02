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
	data   int
	next   *node
	prev   *node
	bottom *node
}

func (n *node) findMiddle() (*node, uint, uint) {
	if n == nil {
		return nil, 0, 0
	}

	var slowIndex, fastIndex uint = 0, 0
	slow := n
	fast := n

	for fast.next != nil {
		slow = slow.next
		fast = fast.next

		slowIndex++
		fastIndex++

		if fast.next != nil {
			fast = fast.next
			fastIndex++
		}
	}

	return slow, slowIndex, fastIndex
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

func (n *node) flattenMixed() *node {
	if n.next != nil {
		n.next.flattenMixed()
	}

	return merge(n, n.next)
}

func merge(a *node, b *node) *node {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	var result *node
	if a.data < b.data {
		result = a
		result.bottom = merge(a.bottom, b)
	} else {
		result = b
		result.bottom = merge(a, b.bottom)
	}

	return result
}

func (n *node) traverseToEndOnBottom() []int {
	values := make([]int, 0)

	curr := n
	for curr != nil {
		values = append(values, curr.data)
		curr = curr.bottom
	}

	return values
}

func (n *node) sortedMerge(head *node) *node {
	if head == nil {
		return n
	}

	var newHead, newTail, currA, currB *node
	if n.data < head.data {
		newHead = n
		newTail = n

		currA = n.next
		currB = head
	} else {
		newHead = head
		newTail = head

		currA = head.next
		currB = n
	}

	for currA != nil || currB != nil {
		var prev *node
		if currA == nil {
			prev = currB
			currB = currB.next
		} else if currB == nil {
			prev = currA
			currA = currA.next
		} else if currA.data < currB.data {
			prev = currA
			currA = currA.next
		} else {
			prev = currB
			currB = currB.next
		}

		newTail.next = prev
		newTail = prev
		newTail.next = nil
	}

	return newHead
}

func (n *node) reverse() *node {
	var prev *node
	curr := n
	next := n.next

	for next != nil {
		curr.next = prev

		prev = curr
		curr = next
		next = next.next
	}

	curr.next = prev
	return curr
}

func (n *node) addAsNumber(head *node) *node {
	if head == nil {
		return n
	}

	currA := n.reverse()
	currB := head.reverse()

	var newHead *node
	carry := 0
	for currA != nil || currB != nil {
		digit := carry

		if currA != nil {
			digit += currA.data
			currA = currA.next
		}
		if currB != nil {
			digit += currB.data
			currB = currB.next
		}

		carry = digit / 10
		digit = digit % 10

		temp := &node{data: digit, next: newHead}
		newHead = temp
	}

	return newHead
}

func (n *node) checkIfPalindrome() bool {
	middle, middleIndex, _ := n.findMiddle()
	secondHead := middle.reverse()
	firstHead := n

	defer secondHead.reverse()

	var index uint = 0
	for ; index < middleIndex; index++ {
		if firstHead.data != secondHead.data {
			return false
		}

		firstHead = firstHead.next
		secondHead = secondHead.next
	}

	return true
}

func (n *node) pushToQueue(val int) {
	curr := n
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &node{data: val}
}

func (n *node) popFromQueue() (int, *node) {
	next := n.next
	n.next = nil

	return n.data, next
}

func (n *node) pushToStack(val int) *node {
	newHead := &node{data: val}
	newHead.next = n
	return newHead
}

func (n *node) popFromStack() (int, *node) {
	return n.popFromQueue()
}
