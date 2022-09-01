package main

import "strings"

func validateUserInput(firstName string, lastName string, age uint, email string, userTickets uint, remainingTickets uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidAge := age >= 18
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidAge, isValidEmail, isValidTicketNumber
}
