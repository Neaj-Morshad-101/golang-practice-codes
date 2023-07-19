package main

import (
	"fmt"
	"strings"
)

func validateUserInput(userData UserData) bool {
	isValidName := len(userData.firstName) >= 2 && len(userData.lastName) >= 2
	isValidEmail := strings.Contains(userData.email, "@")
	isValidNumberOfTickets := userData.numberOfTickets > 0 && userData.numberOfTickets <= remainingTickets

	if !isValidName {
		fmt.Println("first name or last name you entered is too short")
	}
	if !isValidEmail {
		fmt.Println("email address you entered doesn't contain @ sign")
	}
	if !isValidNumberOfTickets {
		fmt.Println("number of tickets you entered is invalid")
	}

	return isValidName && isValidEmail && isValidNumberOfTickets
}
