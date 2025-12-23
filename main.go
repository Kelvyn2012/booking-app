package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go-conference"
	const conferenceTickets uint = 50
	var availableTicket uint = 50
	var bookings []string

	// fmt.Printf("welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", ticketNumber, availableTicket)
	// println("Get your tickets here to attend")

	for {
		greetUsers(conferenceName, availableTicket, conferenceTickets)
		firstName, lastName, email, ticketNumber := getUserDetails()
		isValidEmail, isValidName, isValidTicket := validateUserInput(firstName, lastName, email, ticketNumber, availableTicket)

		if isValidEmail && isValidName && isValidTicket {
			bookings, availableTicket := bookTicket(availableTicket, ticketNumber, bookings, firstName, lastName, conferenceName, email)
			firstNames := getFirstName(bookings)
			fmt.Printf("This are the first names for our bookings: %v\n", firstNames)

			if availableTicket == 0 {
				fmt.Println("The Conference ticket are booked out,come back next year")
				break
			}

		} else {
			if !isValidEmail {
				fmt.Println("Your email does not contain @")
			}
			if !isValidName {
				fmt.Println("Both names should exceed 1 character")
			}
			if !isValidTicket {
				fmt.Println("Correct the number of tickets")
			}
		}

	}

	//fmt.Println(bookings)
	//fmt.Println(len(bookings))
	//fmt.Println(bookings[0])

}
func greetUsers(confName string, ticketNo int, ticketAvail int) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", ticketNo, ticketAvail)
	println("Get your tickets here to attend")
}

func getFirstName(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func validateUserInput(firstName string, lastName string, email string, ticketNumber int, availableTicket int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := ticketNumber <= availableTicket && ticketNumber >= 1
	return isValidEmail, isValidName, isValidTicket

}
func getUserDetails() (string, string, string, int) {
	var firstName string
	var lastName string
	var ticketNumber int
	var email string
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email")
	fmt.Scan(&email)
	fmt.Println("Number of Tickets")
	fmt.Scan(&ticketNumber)
	return firstName, lastName, email, ticketNumber

}
func bookTicket(availableTicket uint, ticketNumber uint, bookings []string, firstName string, lastName string, conferenceName string, email string) ([]string, uint) {
	availableTicket = availableTicket - ticketNumber
	// bookings[0]= firstName+" "+lastName

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for purchasing %v ticket,your confirmation mail will be sent to the email address you provided %v.\n", firstName, lastName, ticketNumber, email)
	fmt.Printf("%v remaining tickets for the %v.\n", availableTicket, conferenceName)
	return bookings, availableTicket

}
