package linkedList

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
