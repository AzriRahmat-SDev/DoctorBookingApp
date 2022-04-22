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
	userSelection = ""
	drVickramSchedule.inOrder()
	fmt.Scanln(&userSelection)
	fmt.Println(userSelection)
	return userSelection
}
