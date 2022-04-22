package main

import (
	"errors"
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

type BST struct {
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

	fmt.Println("\nList of available Doctors: ")
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

func (bst *BST) insertNode(t **binaryNode, time string) error {
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

func (bst *BST) insert(time string) error {
	bst.insertNode(&bst.root, time)
	return nil
}

func (bst *BST) inOrderTraversal(t *binaryNode) error {
	if t != nil {
		bst.inOrderTraversal(t.left)
		fmt.Println(t.time)
		bst.inOrderTraversal(t.right)
	}
	return nil
}

func (bst *BST) inOrder() {
	bst.inOrderTraversal(bst.root)
}

func (bst *BST) findSuccessor(t *binaryNode) *binaryNode {
	for t.right != nil {
		t = t.right
	}
	return t
}

func (bst *BST) removeNode(t **binaryNode, time string) (*binaryNode, error) {
	if *t == nil {
		return nil, errors.New("Tree is empty")
	} else if time < (*t).time {
		(*t).left, _ = bst.removeNode(&(*t).left, time)
	} else if time > (*t).time {
		(*t).right, _ = bst.removeNode(&(*t).right, time)
	} else {
		if (*t).left == nil && (*t).right == nil {
			*t = nil
			return *t, nil
		} else if (*t).left == nil {
			temp := *t
			*t = temp.right
			temp = nil
			return *t, nil
		} else if (*t).right == nil {
			temp := *t
			*t = temp.left
			temp = nil
			return *t, nil
		} else {
			temp := bst.findSuccessor((*t).right)
			(*t) = temp
			(*t).right, _ = bst.removeNode(&(*t).right, time)
		}
	}
	return *t, nil
}

func (bst *BST) remove(time string) {
	bst.root, _ = bst.removeNode(&bst.root, time)
}
