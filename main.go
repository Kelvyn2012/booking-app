package main

import "fmt"

func main() {
	var conferenceName = "Go-conference"
	var availableTicket = 50
	var bookings [50]string

	// fmt.Printf("welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", ticketNumber, availableTicket)
	// println("Get your tickets here to attend")

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
	availableTicket = availableTicket - ticketNumber
	bookings[0] = firstName + " " + lastName

	fmt.Printf("Thank you %v %v for purchasing %v ticket,your confirmation mail will be sent to the email address you provided %v.\n", firstName, lastName, ticketNumber, email)
	fmt.Printf("%v remaining tickets for the %v.\n", availableTicket, conferenceName)
	fmt.Println(bookings)
	fmt.Println(len(bookings))
	fmt.Println(bookings[0])

}
