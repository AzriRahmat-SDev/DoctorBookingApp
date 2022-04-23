package main

import (
	"fmt"
	"strings"
)

func userInputAccess(question string) string {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	return userSelection
}

func userInputRawString(question string) string {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	return userSelection
}

func userInputSelectDoctorName(question string) string {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	removingAllSpaces := strings.ReplaceAll(userSelection, " ", "")
	resultLowerCase := strings.ToLower(removingAllSpaces)
	return resultLowerCase
}

func doctorTimeSlots(question string) string {
	fmt.Println(question)
	drVickramSchedule.inOrder()
	userSelection = ""
	fmt.Println("\nPlease Enter A Slot You Wish To Book(Press q To Exit)")
	fmt.Scanln(&userSelection)
	removingAllSpaces := strings.ReplaceAll(userSelection, " ", "")
	resultLowerCase := strings.ToLower(removingAllSpaces)
	return resultLowerCase
}

func doctorTimeScheduleHandler(question string) string {
	fmt.Println(question)
	doctorRollCall()
	userSelection = ""
	fmt.Scanln(&userSelection)
	return userSelection
}

func selectTimeSlots(question string) string {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	return userSelection
}

func editDoctorScheduleHandler(question string) string {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	return userSelection
}
