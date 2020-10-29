package linkedList

import "testing"

func TestFindMiddle_FromSingleNode(t *testing.T) {
	node := &node{data: 1}
	result := node.findMiddle()
	if node != result {
		t.Errorf("Middle node doesn't match expected node!")
	}
}

func createNodes(count int) []*node {
	nodes := make([]*node, count)

	for index := 0; index < count; index++ {
		nodes[index] = &node{data: index + 1}
	}

	return nodes
}

func createNodesWithSpecificData(data []int) []*node {
	nodes := make([]*node, len(data))

	for index := 0; index < len(data); index++ {
		nodes[index] = &node{data: data[index]}
	}

	return nodes
}

func joinNodes(nodes []*node) {
	for index := 0; index < len(nodes)-1; index++ {
		nodes[index].next = nodes[index+1]
	}
}

func joinNodesWithBottom(nodes []*node) {
	for index := 0; index < len(nodes)-1; index++ {
		nodes[index].bottom = nodes[index+1]
	}
}
func TestFindMiddle_FromOddNodesCount(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)
	result := nodes[0].findMiddle()
	if result != nodes[2] {
		t.Errorf("Middle node doesn't match expected node!")
	}
}

func TestFindMiddle_FromEvenNodesCount(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)
	result := nodes[0].findMiddle()
	if result != nodes[3] {
		t.Errorf("Middle node doesn't match expected node!")
	}
}

func TestRotateCounterClockwise_SingleNodeSingleRotation(t *testing.T) {
	nodes := createNodes(1)
	node := nodes[0]
	result, err := node.rotateCounterClockwise(1)
	if err != nil {
		t.Errorf("Unexpected error = %v", err)
	}
	if result != node {
		t.Errorf("Error = %v", err)
	}
}

func TestRotateCounterClockwise_CountMoreThanNodes(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)
	node := nodes[0]
	_, err := node.rotateCounterClockwise(7)
	if err == nil || err.value != nodeCountExceeded {
		t.Errorf("Expected Error = %s", nodeCountExceeded)
	}
}

func TestRotateCounterClockwise_OddNodesOddRotations(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)
	node := nodes[0]
	result, _ := node.rotateCounterClockwise(3)
	actual := result.traverseToEnd()
	expected := []int{4, 5, 1, 2, 3}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestRotateCounterClockwise_OddNodesEvenRotations(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)
	node := nodes[0]
	result, _ := node.rotateCounterClockwise(4)
	actual := result.traverseToEnd()
	expected := []int{5, 1, 2, 3, 4}
	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestRotateCounterClockwise_EvenNodesOddRotations(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)
	node := nodes[0]
	result, _ := node.rotateCounterClockwise(3)
	actual := result.traverseToEnd()
	expected := []int{4, 5, 6, 1, 2, 3}
	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestRotateCounterClockwise_EvenNodesEvenRotations(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)
	node := nodes[0]
	result, _ := node.rotateCounterClockwise(4)
	actual := result.traverseToEnd()
	expected := []int{5, 6, 1, 2, 3, 4}
	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestReverseInGroups_SingleNodeZeroGroupSize(t *testing.T) {
	nodes := createNodes(1)
	node := nodes[0]
	result := node.reverseInGroups(0)
	if result != node {
		t.Errorf("Expected same node when group size = 0")
	}
}

func TestReverseInGroups_SingleNodeOneGroupSize(t *testing.T) {
	nodes := createNodes(1)
	node := nodes[0]
	result := node.reverseInGroups(1)
	if result != node {
		t.Errorf("Expected same node when group size = 1")
	}
}

func TestReverseInGroups_OddNodesSingleGroupSize(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)
	node := nodes[0]

	result := node.reverseInGroups(1)
	actual := result.traverseToEnd()
	expected := []int{1, 2, 3, 4, 5}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestReverseInGroups_OddNodesOddGroupSize(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)
	node := nodes[0]

	result := node.reverseInGroups(3)
	actual := result.traverseToEnd()
	expected := []int{3, 2, 1, 5, 4}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestReverseInGroups_OddNodesEvenGroupSize(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)
	node := nodes[0]

	result := node.reverseInGroups(4)
	actual := result.traverseToEnd()
	expected := []int{4, 3, 2, 1, 5}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestReverseInGroups_OddNodesExceedingGroupSize(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)
	node := nodes[0]

	result := node.reverseInGroups(8)
	actual := result.traverseToEnd()
	expected := []int{5, 4, 3, 2, 1}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestReverseInGroups_EvenNodesSingleGroupSize(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)
	node := nodes[0]

	result := node.reverseInGroups(1)
	actual := result.traverseToEnd()
	expected := []int{1, 2, 3, 4, 5, 6}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestReverseInGroups_EvenNodesOddGroupSize(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)
	node := nodes[0]

	result := node.reverseInGroups(3)
	actual := result.traverseToEnd()
	expected := []int{3, 2, 1, 6, 5, 4}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestReverseInGroups_EvenNodesEvenGroupSize(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)
	node := nodes[0]

	result := node.reverseInGroups(4)
	actual := result.traverseToEnd()
	expected := []int{4, 3, 2, 1, 6, 5}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestReverseInGroups_EvenNodesExceedingGroupSize(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)
	node := nodes[0]

	result := node.reverseInGroups(8)
	actual := result.traverseToEnd()
	expected := []int{6, 5, 4, 3, 2, 1}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestFindIntersection_SecondListEmpty(t *testing.T) {
	nodes := createNodes(3)
	joinNodes(nodes)
	node := nodes[0]

	result := node.findIntersection(nil)
	if result != nil {
		t.Errorf("Expect nil, but got %v", result)
	}
}

func TestFindIntersection_OddEvenCombination(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)

	headA := nodes[0]
	headB := createNodes(1)[0]
	headB.next = nodes[len(nodes)-1]

	result := headA.findIntersection(headB)
	expected := nodes[len(nodes)-1]
	if result != expected {
		t.Errorf("Expect %v, but got %v", expected, result)
	}
}

func TestFindIntersection_EvenOddCombination(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)

	headA := nodes[0]
	headB := createNodes(1)[0]
	headB.next = nodes[len(nodes)-2]

	result := headA.findIntersection(headB)
	expected := nodes[len(nodes)-2]
	if result != expected {
		t.Errorf("Expect %v, but got %v", expected, result)
	}
}

func TestFindIntersection_BothOddCombination(t *testing.T) {
	nodesA := createNodes(7)
	joinNodes(nodesA)

	nodesB := createNodes(2)
	joinNodes(nodesB)

	nodesB[len(nodesB)-1].next = nodesA[len(nodesA)-3]

	headA := nodesA[0]
	headB := nodesB[0]

	result := headA.findIntersection(headB)
	expected := nodesA[len(nodesA)-3]
	if result != expected {
		t.Errorf("Expect %v, but got %v", expected, result)
	}
}

func TestFindIntersection_BothEvenCombination(t *testing.T) {
	nodesA := createNodes(6)
	joinNodes(nodesA)

	nodesB := createNodes(2)
	joinNodes(nodesB)

	nodesB[len(nodesB)-1].next = nodesA[len(nodesA)-2]

	headA := nodesA[0]
	headB := nodesB[0]

	result := headA.findIntersection(headB)
	expected := nodesA[len(nodesA)-2]
	if result != expected {
		t.Errorf("Expect %v, but got %v", expected, result)
	}
}

func TestFindIntersection_SameNodes(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)

	headA := nodes[0]
	headB := nodes[0]

	result := headA.findIntersection(headB)
	expected := nodes[1]
	if result != expected {
		t.Errorf("Expect %v, but got %v", expected, result)
	}
}

func TestLoopDetect_NodesWithoutLoop(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)

	head := nodes[0]

	result := head.isLoopPresent()
	if result != false {
		t.Errorf("Loop is absent actually!")
	}
}

func TestLoopDetect_LoopToOddNode(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)

	nodes[len(nodes)-1].next = nodes[2]

	head := nodes[0]

	result := head.isLoopPresent()
	if result != true {
		t.Errorf("Loop is present actually!")
	}
}

func TestLoopDetect_LoopToEvenNode(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)

	nodes[len(nodes)-1].next = nodes[3]

	head := nodes[0]

	result := head.isLoopPresent()
	if result != true {
		t.Errorf("Loop is present actually!")
	}
}

func TestLoopRemove_NoLoopNoOperation(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)

	head := nodes[0]
	head.removeLoop()

	actual := head.traverseToEnd()
	expected := []int{1, 2, 3, 4, 5}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestLoopRemove_LoopToOddNode(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)

	nodes[len(nodes)-1].next = nodes[2]

	head := nodes[0]
	head.removeLoop()

	actual := head.traverseToEnd()
	expected := []int{1, 2, 3, 4, 5}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestLoopRemove_LoopToEvenNode(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)

	nodes[len(nodes)-1].next = nodes[3]

	head := nodes[0]
	head.removeLoop()

	actual := head.traverseToEnd()
	expected := []int{1, 2, 3, 4, 5}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestNthFromLast_SingleNode(t *testing.T) {
	nodes := createNodes(1)
	joinNodes(nodes)

	head := nodes[0]
	result := head.getNthFromLast(0)
	expected := 1

	if *result != expected {
		t.Errorf("Expected = %v, actual = %v", expected, result)
	}
}

func TestNthFromLast_ExceedsNodeCount(t *testing.T) {
	nodes := createNodes(5)
	joinNodes(nodes)

	head := nodes[0]
	result := head.getNthFromLast(6)
	var expected *int = nil

	if result != expected {
		t.Errorf("Expected = %v, actual = %v", expected, result)
	}
}

func TestNthFromLast_MiddleValue(t *testing.T) {
	nodes := createNodes(6)
	joinNodes(nodes)

	head := nodes[0]
	result := head.getNthFromLast(2)
	expected := 4

	if *result != expected {
		t.Errorf("Expected = %v, actual = %v", expected, result)
	}
}

func TestFlattening_SingleNode(t *testing.T) {
	nodes := createNodes(1)
	head := nodes[0]
	actual := head.traverseToEndOnBottom()
	expected := []int{1}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestFlattening_MultipleLevels(t *testing.T) {
	first := createNodesWithSpecificData([]int{5, 7, 8, 30})
	joinNodesWithBottom(first)

	second := createNodesWithSpecificData([]int{10, 20})
	joinNodesWithBottom(second)

	third := createNodesWithSpecificData([]int{19, 22, 50})
	joinNodesWithBottom(third)

	fourth := createNodesWithSpecificData([]int{28, 35, 40, 45})
	joinNodesWithBottom(fourth)

	joinNodes([]*node{first[0], second[0], third[0], fourth[0]})

	head := first[0]
	head = head.flattenMixed()

	actual := head.traverseToEndOnBottom()
	expected := []int{5, 7, 8, 10, 19, 20, 22, 28, 30, 35, 40, 45, 50}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestSortedMerge_EmptyHead(t *testing.T) {
	first := &node{data: 1}

	head := first.sortedMerge(nil)
	actual := head.traverseToEnd()
	expected := []int{1}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestSortedMerge_FirstSmallerThanSecond(t *testing.T) {
	first := createNodesWithSpecificData([]int{5, 7, 8, 30})
	joinNodes(first)

	second := createNodesWithSpecificData([]int{10, 20})
	joinNodes(second)

	head := first[0].sortedMerge(second[0])
	actual := head.traverseToEnd()
	expected := []int{5, 7, 8, 10, 20, 30}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestSortedMerge_FirstLargerThanSecond(t *testing.T) {
	first := createNodesWithSpecificData([]int{5, 30})
	joinNodes(first)

	second := createNodesWithSpecificData([]int{7, 10, 15, 20, 40})
	joinNodes(second)

	head := first[0].sortedMerge(second[0])
	actual := head.traverseToEnd()
	expected := []int{5, 7, 10, 15, 20, 30, 40}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}

func TestSortedMerge_FirstLesserThanSecond(t *testing.T) {
	first := createNodesWithSpecificData([]int{7, 30, 40})
	joinNodes(first)

	second := createNodesWithSpecificData([]int{5, 35, 35})
	joinNodes(second)

	head := first[0].sortedMerge(second[0])
	actual := head.traverseToEnd()
	expected := []int{5, 7, 30, 35, 35, 40}

	if len(actual) != len(expected) {
		t.Errorf("Actual result size %d != Expected result size %d", len(actual), len(expected))
		return
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			return
		}
	}
}
