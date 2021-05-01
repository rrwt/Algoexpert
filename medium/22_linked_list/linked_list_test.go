package linked_list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	linkedList := NewDoublyLinkedList()
	one := NewNode(1)
	two := NewNode(2)
	three := NewNode(3)
	three2 := NewNode(3)
	three3 := NewNode(3)
	four := NewNode(4)
	five := NewNode(5)
	six := NewNode(6)
	bindNodes(one, two)
	bindNodes(two, three)
	bindNodes(three, four)
	bindNodes(four, five)
	linkedList.Head = one
	linkedList.Tail = five

	linkedList.SetHead(four)
	require.Equal(t, []int{4, 1, 2, 3, 5}, getNodeValuesHeadToTail(linkedList))
	require.Equal(t, []int{5, 3, 2, 1, 4}, getNodeValuesTailToHead(linkedList))

	linkedList.SetTail(six)
	require.Equal(t, []int{4, 1, 2, 3, 5, 6}, getNodeValuesHeadToTail(linkedList))
	require.Equal(t, []int{6, 5, 3, 2, 1, 4}, getNodeValuesTailToHead(linkedList))

	linkedList.InsertBefore(six, three)
	require.Equal(t, []int{4, 1, 2, 5, 3, 6}, getNodeValuesHeadToTail(linkedList))
	require.Equal(t, []int{6, 3, 5, 2, 1, 4}, getNodeValuesTailToHead(linkedList))

	linkedList.InsertAfter(six, three2)
	require.Equal(t, []int{4, 1, 2, 5, 3, 6, 3}, getNodeValuesHeadToTail(linkedList))
	require.Equal(t, []int{3, 6, 3, 5, 2, 1, 4}, getNodeValuesTailToHead(linkedList))

	linkedList.InsertAtPosition(1, three3)
	require.Equal(t, []int{3, 4, 1, 2, 5, 3, 6, 3}, getNodeValuesHeadToTail(linkedList))
	require.Equal(t, []int{3, 6, 3, 5, 2, 1, 4, 3}, getNodeValuesTailToHead(linkedList))

	linkedList.RemoveNodesWithValue(3)
	require.Equal(t, []int{4, 1, 2, 5, 6}, getNodeValuesHeadToTail(linkedList))
	require.Equal(t, []int{6, 5, 2, 1, 4}, getNodeValuesTailToHead(linkedList))

	linkedList.Remove(two)
	require.Equal(t, []int{4, 1, 5, 6}, getNodeValuesHeadToTail(linkedList))
	require.Equal(t, []int{6, 5, 1, 4}, getNodeValuesTailToHead(linkedList))

	require.Equal(t, true, linkedList.ContainsNodeWithValue(5))
}

func NewNode(value int) *Node { return &Node{Value: value} }

func getNodeValuesHeadToTail(ll *DoublyLinkedList) []int {
	values := []int{}
	node := ll.Head
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	return values
}

func getNodeValuesTailToHead(ll *DoublyLinkedList) []int {
	values := []int{}
	node := ll.Tail
	for node != nil {
		values = append(values, node.Value)
		node = node.Prev
	}
	return values
}

func bindNodes(nodeOne *Node, nodeTwo *Node) {
	nodeOne.Next = nodeTwo
	nodeTwo.Prev = nodeOne
}