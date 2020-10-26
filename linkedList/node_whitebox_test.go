package linkedList

import "testing"

func TestFindMiddle_FromSingleNode(t *testing.T) {
	node := &node{data: 1}
	result := node.findMiddle()
	if node != result {
		t.Errorf("FindMiddle_SingleNode: Middle node doesn't match expected node!")
	}
}

func TestFindMiddle_FromOddNodesCount(t *testing.T) {
	nodes := []*node{
		&node{data: 1}, &node{data: 2}, &node{data: 3},
		&node{data: 4}, &node{data: 5},
	}
	for index := 0; index < len(nodes)-1; index++ {
		nodes[index].next = nodes[index+1]
	}
	result := nodes[0].findMiddle()
	if result != nodes[2] {
		t.Errorf("FindMiddle_OddNodesCount: Middle node doesn't match expected node!")
	}
}

func TestFindMiddle_FromEvenNodesCount(t *testing.T) {
	nodes := []*node{
		&node{data: 1}, &node{data: 2}, &node{data: 3},
		&node{data: 4}, &node{data: 5}, &node{data: 6},
	}
	for index := 0; index < len(nodes)-1; index++ {
		nodes[index].next = nodes[index+1]
	}
	result := nodes[0].findMiddle()
	if result != nodes[3] {
		t.Errorf("FindMiddle_EvenNodesCount: Middle node doesn't match expected node!")
	}
}
