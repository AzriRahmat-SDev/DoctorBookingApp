package main

import (
	"errors"
	"fmt"
)

func showMainMenu() {
	fmt.Println("\n1. Appointment Booking")
	fmt.Println("2. Admin Access")
	fmt.Scanln(&userInputMainMenu)
}

func featureMenu() {
	fmt.Println("\nMain Menu")
	fmt.Println("===========")
	fmt.Println("1. View Available Doctors")
	fmt.Println("2. Make An Appointment")
	fmt.Println("3. Search Doctors By Time Slots")
	fmt.Println("4. Edit Appointment")
	fmt.Println("5. Exit Program")
	fmt.Println("\nPlease Choose An Option:")
	fmt.Scanln(&userInputFeatureMenu)
}

//Option 1
func doctorRollCall() { //should be change to current booking that you have
	doctorLists.printAllDoctorNodes()
}

//Option 2
func bookingAppointment() {
	doctorName := userInputSelectDoctorName("\nPlease select the doctor you would want to book:")

	if doctorName == "dr.vickram" {
		removingTimeSlotsVickram()
	} else if doctorName == "dr.fazuli" {
		removingTimeSlotsFazuli()
	} else if doctorName == "dr.idris" {
		removingTimeSlotsIdris()
	} else {
		panic(errors.New("Doctor name entered is not in the roster"))
	}
}

//Option 2
func removingTimeSlotsVickram() {
	timeSlot := doctorTimeSlots("\nThese are available Timeslots")
	removedTimeSlot := drVickramSchedule.remove(timeSlot)

	switch removedTimeSlot.time {
	case "slot1":
		vickramBookedSlots[0] = removedTimeSlot
	case "slot2":
		vickramBookedSlots[1] = removedTimeSlot
	case "slot3":
		vickramBookedSlots[2] = removedTimeSlot
	case "slot4":
		vickramBookedSlots[3] = removedTimeSlot
	case "slot5":
		vickramBookedSlots[4] = removedTimeSlot
	case "slot6":
		vickramBookedSlots[5] = removedTimeSlot
	case "slot7":
		vickramBookedSlots[6] = removedTimeSlot
	case "slot8":
		vickramBookedSlots[7] = removedTimeSlot
	case "slot9":
		vickramBookedSlots[8] = removedTimeSlot
	case "q":
		main()
	default:
		panic(errors.New("Invalid time slot"))
	}
	main()
}

func removingTimeSlotsFazuli() {
	timeSlot := doctorTimeSlots("\nThese are available Timeslots")
	removedTimeSlot := drFazuliSchedule.remove(timeSlot)

	switch removedTimeSlot.time {
	case "slot1":
		fazuliBookedSlots[0] = removedTimeSlot
	case "slot2":
		fazuliBookedSlots[1] = removedTimeSlot
	case "slot3":
		fazuliBookedSlots[2] = removedTimeSlot
	case "slot4":
		fazuliBookedSlots[3] = removedTimeSlot
	case "slot5":
		fazuliBookedSlots[4] = removedTimeSlot
	case "slot6":
		fazuliBookedSlots[5] = removedTimeSlot
	case "slot7":
		fazuliBookedSlots[6] = removedTimeSlot
	case "slot8":
		fazuliBookedSlots[7] = removedTimeSlot
	case "slot9":
		fazuliBookedSlots[8] = removedTimeSlot
	case "q":
		main()
	default:
		panic(errors.New("Invalid time slot"))
	}
	main()
}

func removingTimeSlotsIdris() {
	timeSlot := doctorTimeSlots("\nThese are available Timeslots")
	removedTimeSlot := drIdrisSchedule.remove(timeSlot)

	switch removedTimeSlot.time {
	case "slot1":
		idrisBookedSlots[0] = removedTimeSlot
	case "slot2":
		idrisBookedSlots[1] = removedTimeSlot
	case "slot3":
		idrisBookedSlots[2] = removedTimeSlot
	case "slot4":
		idrisBookedSlots[3] = removedTimeSlot
	case "slot5":
		idrisBookedSlots[4] = removedTimeSlot
	case "slot6":
		idrisBookedSlots[5] = removedTimeSlot
	case "slot7":
		idrisBookedSlots[6] = removedTimeSlot
	case "slot8":
		idrisBookedSlots[7] = removedTimeSlot
	case "slot9":
		idrisBookedSlots[8] = removedTimeSlot
	case "q":
		main()
	default:
		panic(errors.New("Invalid time slot"))
	}
	main()
}

//Option 3
func searchTimeSlots() {
	doctorTimeSchedule := doctorTimeScheduleHandler("\nPlease Select The Available Doctors")
	if doctorTimeSchedule == "dr.vickram" {
		selectedTimeByUser := selectTimeSlots("\nPlease Select A Time (Slot1-Slot9)")
		//fmt.Println(vickramBookedSlots) //why the position is not at the right index??(this is because of the removeNode() in the BST)
		for i := range vickramBookedSlots {
			if vickramBookedSlots[i] == nil {
				continue
			} else if selectedTimeByUser == vickramBookedSlots[i].time {
				fmt.Println("Yes This Slot Is Available For The Selected Doctor")
				return
			} else {
				fmt.Println("Sorry This Slot Is Not Available For The Selected Doctor")
				return
			}
		}
	} else if doctorTimeSchedule == "dr.fazuli" {
		selectedTimeByUser := selectTimeSlots("\nPlease Select A Time (Slot1-Slot9)")
		//fmt.Println(fazuliBookedSlots) //why the position is not at the right index??(this is because of the removeNode() in the BST)
		for i := range fazuliBookedSlots {
			if fazuliBookedSlots[i] == nil {
				continue
			} else if selectedTimeByUser == fazuliBookedSlots[i].time {
				fmt.Println("Yes This Slot Is Available For The Selected Doctor")
				return
			} else {
				fmt.Println("Sorry This Slot Is Not Available For The Selected Doctor")
				return
			}
		}
	} else if doctorTimeSchedule == "dr.idris" {
		selectedTimeByUser := selectTimeSlots("\nPlease Select A Time (Slot1-Slot9)")
		//fmt.Println(idrisBookedSlots) //why the position is not at the right index??(this is because of the removeNode() in the BST)
		for i := range idrisBookedSlots {
			if idrisBookedSlots[i] == nil {
				continue
			} else if selectedTimeByUser == idrisBookedSlots[i].time {
				fmt.Println("Yes This Slot Is Available For The Selected Doctor")
				return
			} else {
				fmt.Println("Sorry This Slot Is Not Available For The Selected Doctor")
				return
			}
		}
	}
	featureMenu()
}

//Option 4
func editAppointment() {
	editDoctorSchedule := editDoctorScheduleHandler("\nPlease Select The Available Doctors")
	if editDoctorSchedule == "dr.vickram" {
		editTimeByUser := selectTimeSlots("\nPlease Select A Time (Slot1-Slot9)")
		for i := range vickramBookedSlots {
			if vickramBookedSlots[i].time == "" {
				continue
			} else if editTimeByUser == vickramBookedSlots[i].time {
				vickramBookedSlots[i] = nil
				fmt.Println("\nThis Slot Is Open To The Public Now")
				return
			} else {
				fmt.Println("\nSorry This Slot Has Not Been Booked")
				return
			}
		}
		drVickramSchedule.insert(editTimeByUser)
	}
}
