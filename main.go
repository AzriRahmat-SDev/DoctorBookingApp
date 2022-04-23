package main

import (
	"fmt"
)

var doctorLists = &doctorsLinkedList{nil, 0}
var drVickramSchedule = &BST{nil}
var drFazuliSchedule = &BST{nil}
var drIdrisSchedule = &BST{nil}
var userInputMainMenu int
var userInputFeatureMenu int
var userInputBooking string
var userInputRemove string

//Declaring Array for Timeslots
var bookedSlotsList = [][9]*binaryNode{vickramBookedSlots, fazuliBookedSlots, idrisBookedSlots}

var vickramBookedSlots = [9]*binaryNode{}
var fazuliBookedSlots = [9]*binaryNode{}
var idrisBookedSlots = [9]*binaryNode{}

var userSelection string

func init() {
	doctorLists.addDoctors("Dr.Vickram")
	doctorLists.addDoctors("Dr.Fazuli")
	doctorLists.addDoctors("Dr.Idris")
	doctorLists.addDoctors("Dr.Sim")

	//Declaring Open schedule for the doctors
	//Vickram
	drVickramSchedule.insert("Slot4")
	drVickramSchedule.insert("Slot3")
	drVickramSchedule.insert("Slot6")
	drVickramSchedule.insert("Slot8")
	drVickramSchedule.insert("Slot9")

	//Fazuli
	drFazuliSchedule.insert("Slot4")
	drFazuliSchedule.insert("Slot2")
	drFazuliSchedule.insert("Slot6")
	drFazuliSchedule.insert("Slot3")
	drFazuliSchedule.insert("Slot9")
	drFazuliSchedule.insert("Slot7")
	drFazuliSchedule.insert("Slot8")

	//Idris
	drIdrisSchedule.insert("Slot5")
	drIdrisSchedule.insert("Slot1")
	drIdrisSchedule.insert("Slot3")
	drIdrisSchedule.insert("Slot6")
	drIdrisSchedule.insert("Slot9")
	drIdrisSchedule.insert("Slot7")

	//Declaring closed schedule for Doctors
	//vickram
	drVickramSchedule.insert("Slot1")
	drVickramSchedule.insert("Slot2")
	drVickramSchedule.insert("Slot5")
	drVickramSchedule.insert("Slot7")
	initVTimeSlots1 := drVickramSchedule.remove("Slot1")
	vickramBookedSlots[0] = initVTimeSlots1
	initVTimeSlots2 := drVickramSchedule.remove("Slot2")
	vickramBookedSlots[1] = initVTimeSlots2
	initVTimeSlots4 := drVickramSchedule.remove("Slot3")
	vickramBookedSlots[4] = initVTimeSlots4
	initVTimeSlots7 := drVickramSchedule.remove("Slot1")
	vickramBookedSlots[6] = initVTimeSlots7

	//Fazuli
	drFazuliSchedule.insert("Slot1")
	drFazuliSchedule.insert("Slot5")
	initFTimeSlots2 := drFazuliSchedule.remove("Slot2")
	fazuliBookedSlots[1] = initFTimeSlots2
	initFTimeSlots5 := drFazuliSchedule.remove("Slot5")
	fazuliBookedSlots[4] = initFTimeSlots5

	//Idris
	drIdrisSchedule.insert("Slot2")
	drIdrisSchedule.insert("Slot4")
	drIdrisSchedule.insert("Slot8")
	initITimeSlots2 := drIdrisSchedule.remove("Slot2")
	fazuliBookedSlots[1] = initITimeSlots2
	initITimeSlots4 := drIdrisSchedule.remove("Slot4")
	fazuliBookedSlots[3] = initITimeSlots4
	initITimeSlots8 := drIdrisSchedule.remove("Slot8")
	fazuliBookedSlots[7] = initITimeSlots8
}

func main() {
	fmt.Println("Welcome to Doctor Appointment Booking System.")
	fmt.Println("=============================================")
	showMainMenu()

	if userInputMainMenu == 1 {
		featureMenu()
	} else if userInputMainMenu == 2 {
		showMainMenu()
	}

	if userInputFeatureMenu == 1 {
		doctorRollCall()
	} else if userInputFeatureMenu == 2 {
		doctorRollCall()
		bookingAppointment()
	} else if userInputFeatureMenu == 3 {
		searchTimeSlots()
	}
}
