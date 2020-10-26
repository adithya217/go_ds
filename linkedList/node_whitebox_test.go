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

func joinNodes(nodes []*node) {
	for index := 0; index < len(nodes)-1; index++ {
		nodes[index].next = nodes[index+1]
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
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
	}
	for index := 0; index < len(expected); index++ {
		actualNum := actual[index]
		expectedNum := expected[index]
		if actualNum != expectedNum {
			t.Errorf("Actual data %d != Expected data %d", actualNum, expectedNum)
			break
		}
	}
}
