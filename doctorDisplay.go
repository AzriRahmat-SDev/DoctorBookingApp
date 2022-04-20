package main

import (
	"fmt"
)

type doctorNode struct {
	name string
	next *doctorNode
}

type doctorLinkedList struct {
	head *doctorNode
	size int
}

func (r *doctorLinkedList) addDoctors(name string) error {
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

func (r *doctorLinkedList) printAllDoctorNodes() error {
	if r.head == nil {
		fmt.Println("There are no doctors in the list")
		return nil
	}
	currentNode := r.head
	for currentNode != nil {
		currentNode := currentNode.next
		fmt.Println(currentNode.name)
	}
	return nil
}
