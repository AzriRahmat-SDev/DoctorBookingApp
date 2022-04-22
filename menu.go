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
	fmt.Println("3. Search Doctors By Time Slots")
	fmt.Println("4. Delete Appointment")
	fmt.Println("5. Exit Program")
	fmt.Println("\nPlease Choose An Option:")
	fmt.Scanln(&userInputFeatureMenu)
}

func doctorRollCall() { //should be change to current booking that you have
	doctorLists.printAllDoctorNodes()
}

//Option 1
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

//Option 2
func removingTimeSlots() {
	timeSlot := doctorTimeSlots("\nThese are available Timeslots")
	fmt.Println("this is difficlt man!", timeSlot)
	removedTimeSlot := drVickramSchedule.remove(timeSlot)

	switch removedTimeSlot.time {
	case "Slot1":
		vickramBookedSlots[0] = removedTimeSlot
	case "Slot2":
		vickramBookedSlots[1] = removedTimeSlot
	case "Slot3":
		vickramBookedSlots[2] = removedTimeSlot
	case "Slot4":
		vickramBookedSlots[3] = removedTimeSlot
	case "Slot5":
		vickramBookedSlots[4] = removedTimeSlot
	case "Slot6":
		vickramBookedSlots[5] = removedTimeSlot
	case "Slot7":
		vickramBookedSlots[6] = removedTimeSlot
	case "Slot8":
		vickramBookedSlots[7] = removedTimeSlot
	case "Slot9":
		vickramBookedSlots[8] = removedTimeSlot
	}

	fmt.Println("\nShowing Current Schedule After Booking")
	if timeSlot == "q" {
		main()
	}
	removingTimeSlots()
}

//Option 3
func searchTimeSlots() {
	doctorTimeSchedule := doctorTimeScheduleHandler("\nPlease Select The Available Doctors")
	if doctorTimeSchedule == "dr.vickram" {
		selectedTimeByUser := selectTimeSlots("\nPlease Select A Time (Slot1-Slot9)")
		//fmt.Println(vickramBookedSlots) //why the position is not at the right index??(this is because of the removeNode() in the BST)

		for i := range vickramBookedSlots {
			if vickramBookedSlots[i] == nil {
				// fmt.Println("Work please")
				continue
			} else if selectedTimeByUser == vickramBookedSlots[i].time {
				fmt.Println("Yes This Slot is available")
			} else {
				fmt.Println("Sorry This Slot is not available")
			}
		}

		// for i := 0; i < len(vickramBookedSlots)-1; i++ {
		// 	if selectedTimeByUser == vickramBookedSlots[i].time {
		// 		fmt.Println("Yes This Slot is available")
		// 	} else {
		// 		fmt.Println("Sorry This Slot is not available")
		// 	}
		// }
	}
}
