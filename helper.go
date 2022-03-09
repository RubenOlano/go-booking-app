package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 1 && len(lastName) > 1
	isValidEmail := strings.Contains(email, "@")
	isValidNum := userTickets <= remainingTickets && userTickets > 0

	return isValidName, isValidEmail, isValidNum
}
