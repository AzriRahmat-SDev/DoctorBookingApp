package main

import (
	"fmt"
)

var doctorLists = &doctorsLinkedList{nil, 0}
var drVickramSchedule = &bst{nil}
var drFazuliSchedule = &bst{nil}
var drIdrisSchedule = &bst{nil}
var drSimSchedule = &bst{nil}
var userInputMainMenu int
var userInputBooking string

func init() {
	doctorLists.addDoctors("Dr.Vickram")
	doctorLists.addDoctors("Dr.Fazuli")
	doctorLists.addDoctors("Dr.Idris")
	doctorLists.addDoctors("Dr.Sim")

	//Declaring Open schedule for the doctors

	//Vickram
	drVickramSchedule.insert("Slot 4")
	drVickramSchedule.insert("Slot 3")
	drVickramSchedule.insert("Slot 5")
	drVickramSchedule.insert("Slot 8")
	drVickramSchedule.insert("Slot 1")

	//Fazuli
	drFazuliSchedule.insert("Slot 4")
	drFazuliSchedule.insert("Slot 2")
	drFazuliSchedule.insert("Slot 6")
	drFazuliSchedule.insert("Slot 3")
	drFazuliSchedule.insert("Slot 9")
	drFazuliSchedule.insert("Slot 7")
	drFazuliSchedule.insert("Slot 8")

	//Idris
	drIdrisSchedule.insert("Slot 5")
	drIdrisSchedule.insert("Slot 1")
	drIdrisSchedule.insert("Slot 3")
	drIdrisSchedule.insert("Slot 6")
	drIdrisSchedule.insert("Slot 9")
	drIdrisSchedule.insert("Slot 7")

	//Sim
	drSimSchedule.insert("Slot 3")
	drSimSchedule.insert("Slot 1")
	drSimSchedule.insert("Slot 6")
	drSimSchedule.insert("Slot 8")
}

func main() {
	fmt.Println("Welcome to Doctor Appointment Booking System.")
	showMainMenu()
	fmt.Scanln(&userInputMainMenu)
	if userInputMainMenu == 1 {
		doctorLists.printAllDoctorNodes()
		bookingMenu()
	} else if userInputMainMenu == 2 {
		fmt.Println()
	}
}

/*
okay i figured out the issue.
the name you gave your variable cannot share the same name as the type you declared
this no work ==> doctorsLinkedList := &doctorsLinkedList
this works ==> doctorLists = &doctorsLinkedList{nil, 0}

i think its because the compiler thinks youre trying to re-declaring the type to something else
*/
