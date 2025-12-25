package main

import "strings"

func validateUserInput(userTickets uint,firstName string,lastName string, email string)(bool,bool,bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail,isValidName,isValidUserTickets
}