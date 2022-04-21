package main

import (
	"fmt"
)

type doctorNode struct {
	name string
	next *doctorNode
}
type doctorsLinkedList struct {
	head *doctorNode
	size int
}

type binaryNode struct {
	time  string
	left  *binaryNode
	right *binaryNode
}

type bst struct {
	root *binaryNode
}

func (r *doctorsLinkedList) addDoctors(name string) error {
	newNode := &doctorNode{
		name: name,
		next: nil,
	}
	if r.head == nil {
		r.head = newNode
	} else {
		currentNode := r.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	r.size++
	return nil
}

func (r *doctorsLinkedList) printAllDoctorNodes() error {

	fmt.Println("List of available Doctors: ")
	if r.head == nil {
		fmt.Println("There are no doctors in the list")
		return nil
	}
	currentNode := r.head
	fmt.Println(currentNode.name)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Println(currentNode.name)
	}
	return nil
}

func (bst *bst) insertNode(t **binaryNode, time string) error {
	if *t == nil {
		newNode := &binaryNode{
			time:  time,
			left:  nil,
			right: nil,
		}
		*t = newNode
		return nil
	}
	if time < (*t).time {
		bst.insertNode(&(*t).left, time)
	} else {
		bst.insertNode(&(*t).right, time)
	}
	return nil
}
func (bst *bst) insert(time string) error {
	bst.insertNode(&bst.root, time)
	return nil
}

func (bst *bst) inOrderTraversal(t *binaryNode) error {
	if t != nil {
		bst.inOrderTraversal(t.left)
		fmt.Println(t.time)
		bst.inOrderTraversal(t.right)
	}
	return nil
}

func (bst *bst) inOrder() {
	bst.inOrderTraversal(bst.root)
}
