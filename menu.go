package main

import (
	"fmt"
	"strings"
)

func showMainMenu() {
	fmt.Println("For Booking for an appointment, Press 1")
	fmt.Println("For access of admin rights, Press 2")

}

func bookingMenu() {
	fmt.Println("Please select the doctor you would want to book:")
	fmt.Scanln(&userInputBooking)
	removingAllSpaces := strings.ReplaceAll(userInputBooking, " ", "")
	result := strings.ToLower(removingAllSpaces)

	if result == "dr.vickram" {
		fmt.Println("These are the available slots")
		drVickramSchedule.inOrder()
		fmt.Println("Please enter the timing you want to book (I.e: Slot 1)")
		fmt.Scanln(userInputRemove)
		drVickramSchedule.remove(userInputRemove)
	}

	if result == "dr.fazuli" {
		drFazuliSchedule.inOrder()
		fmt.Scanln()
	}

	if result == "dr.idris" {
		drIdrisSchedule.inOrder()
		fmt.Scanln()
	}

	if result == "dr.simanntan" {
		drSimSchedule.inOrder()
		fmt.Scanln()
	}
}
