package main

import "fmt"

type heap struct {
	data []int
	root int
	last int
}

func (h *heap) delete() int {
	popped := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.trickleDown(0)
	return popped
}

func (h *heap) trickleDown(index int) {
	leftChildIndex := getLeftChildIndex(index)
	rightChildIndex := getRightChildIndex(index)

	greaterIndex := h.findIndexOfGreaterValue(leftChildIndex, rightChildIndex)

}

func (h *heap) findIndexOfGreaterValue(leftChildIndex, rightChildIndex int) int {
	var leftValue, rightValue int
	if leftChildIndex < len(h.data) {
		leftValue = h.data[leftChildIndex]
	}
	if rightChildIndex < len(h.data) {
		rightValue = h.data[rightChildIndex]
	}
	if leftValue > rightValue {
		return leftChildIndex
	} else if rightValue > leftValue {
		return rightChildIndex
	}
}

func (h *heap) insert(i int) {
	h.data = append(h.data, i)
	newNodeIndex := len(h.data) - 1

	parentIndex := getParentIndex(newNodeIndex)
	h.trickleUp(newNodeIndex, parentIndex)
}

func (h *heap) trickleUp(childIndex, parentIndex int) {
	if h.data[childIndex] > h.data[parentIndex] {
		h.data[parentIndex], h.data[childIndex] = h.data[childIndex], h.data[parentIndex]
		h.trickleUp(parentIndex, getParentIndex(parentIndex))
	}
}

func getLeftChildIndex(index int) int {
	return (index * 2) + 1
}

func getRightChildIndex(index int) int {
	return (index * 2) + 2
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func (h *heap) insertValues(slice []int) {
	for _, value := range slice {
		h.insert(value)
	}
}

func main() {
	myHeap := heap{}
	myHeap.insertValues([]int{2, 3, 8, 12, 15, 16, 25, 50, 86, 87, 88, 100, 120})
	fmt.Println(myHeap.data)
	myHeap.delete()
	fmt.Println(myHeap.data)
}
