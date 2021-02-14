package main

import "fmt"

type linkedList struct {
	Head   *node
	Length int
}

type node struct {
	data int
	next *node
}

func (l *linkedList) prepend(node *node) {
	second := l.Head
	l.Head = node
	l.Head.next = second
	l.Length++
}

func (l linkedList) printData() {
	node := l.Head
	for l.Length != 0 {
		fmt.Println(node.data)
		node = node.next
		l.Length--
	}
}

func (l linkedList) read(index int) {
	currentNode := l.Head
	currentIndex := 0
	for currentIndex < index {
		currentNode = currentNode.next
		currentIndex++
	}
	fmt.Println(currentNode.data)
}

func (l linkedList) indexOf(value int) {
	currentNode := l.Head
	currentValue := currentNode.data
	currentIndex := 0
	for currentValue != value && currentIndex <= l.Length {
		currentNode = currentNode.next
		currentValue = currentNode.data
		currentIndex++
	}
	fmt.Println(currentIndex)
}

func (l *linkedList) deleteAtIndex(index int) {
	if index == 0 {
		l.Head = l.Head.next
	} else {
		currentNode := l.Head
		nextNode := l.Head.next
		currentIndex := 1
		for currentIndex != index {
			currentNode = currentNode.next
			nextNode = currentNode.next
			currentIndex++
		}
		currentNode.next = nextNode.next
	}
	l.Length--
}

func (l *linkedList) insertAtIndex(index, value int) {
	newNode := &node{data: value}
	if index == 0 {
		l.prepend(newNode)
	} else {
		previousNode := l.Head
		currentNode := l.Head.next
		currentIndex := 1
		for currentIndex != index {
			previousNode = currentNode
			currentNode = previousNode.next
			currentIndex++
		}
		previousNode.next = newNode
		newNode.next = currentNode
		l.Length++
	}
}

func main() {
	node1 := &node{data: 2}
	node2 := &node{data: 3}
	node3 := &node{data: 4}
	node4 := &node{data: 5}

	myList := linkedList{}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	myList.printData()
	fmt.Printf("\n")
	myList.insertAtIndex(3, 53)
	myList.printData()
	myList.deleteAtIndex(0)
	fmt.Printf("\n")
	myList.printData()
}
