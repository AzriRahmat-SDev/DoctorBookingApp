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
var vickramBookedSlots = [9]*binaryNode{}
var fazuliBookedSlots = [9]*binaryNode{}
var idrisBookedSlots = [9]*binaryNode{}

var userSelection string

func init() {
	doctorLists.addDoctors("Dr.Vickram")
	doctorLists.addDoctors("Dr.Fazuli")
	doctorLists.addDoctors("Dr.Idris")

	//Declaring Open schedule for the doctors
	//Vickram
	drVickramSchedule.insert("Slot3")
	drVickramSchedule.insert("Slot6")
	drVickramSchedule.insert("Slot8")
	drVickramSchedule.insert("Slot9")
	drVickramSchedule.insert("Slot4")

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
	initVTimeSlots4 := drVickramSchedule.remove("Slot5")
	vickramBookedSlots[4] = initVTimeSlots4
	initVTimeSlots7 := drVickramSchedule.remove("Slot7")
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
	fmt.Println("\nWelcome to Doctor Appointment Booking System.")
	fmt.Println("=============================================")
	showMainMenu()

	if userInputMainMenu == 1 {
		featureMenu()
	} else {
		fmt.Println("Invalid Option")
		main()
	}

	//Feature Menu
	if userInputFeatureMenu == 1 {
		doctorRollCall()
	} else if userInputFeatureMenu == 2 {
		doctorRollCall()
		bookingAppointment()
		main()
	} else if userInputFeatureMenu == 3 {
		searchTimeSlots()
		main()
	} else if userInputFeatureMenu == 4 {
		editAppointment()
		main()
	} else if userInputFeatureMenu == 5 {
		main()
	}
}
