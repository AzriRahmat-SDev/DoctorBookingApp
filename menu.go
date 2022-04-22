package main

import (
	"fmt"
)

func showMainMenu() {
	fmt.Println("\n1. Appointment Booking")
	fmt.Println("2. Admin Access")
	fmt.Scanln(&userInputMainMenu)
}

func featureMenu() {
	fmt.Println("Main Menu")
	fmt.Println("===========")
	fmt.Println("1. View Available Doctors")
	fmt.Println("2. Make An Appointment")
	fmt.Println("3. Reschedule Your Appointment")
	fmt.Println("4. Delete Appointment")
	fmt.Println("5. Exit Program")
	fmt.Println("\nPlease Choose An Option:")
	fmt.Scanln(&userInputFeatureMenu)
}

func doctorRollCall() { //should be change to current booking that you have
	doctorLists.printAllDoctorNodes()
}

func bookingAppointment() {
	doctorName := userInputSelectDoctorName("\nPlease select the doctor you would want to book:")

	if doctorName == "dr.vickram" {
		removingTimeSlots()
	}

	if doctorName == "dr.fazuli" {
		fmt.Scanln()
	}

	if doctorName == "dr.idris" {
		drIdrisSchedule.inOrder()
		fmt.Scanln()
	}

	if doctorName == "dr.simanntan" {
		drSimSchedule.inOrder()
		fmt.Scanln()
	}
}

func removingTimeSlots() {
	timeSlot := doctorTimeSlots("\nThese are available Timeslots")
	drVickramSchedule.remove(timeSlot)

	drVickramSchedule.remove("Slot 1")
	fmt.Println("\nShowing Current Schedule After Booking")
	drVickramSchedule.inOrder()
}
