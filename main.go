package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go-conference"
	var availableTicket = 50
	var bookings []string

	// fmt.Printf("welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", ticketNumber, availableTicket)
	// println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var ticketNumber int
	var email string

	for {
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your Email")
		fmt.Scan(&email)
		fmt.Println("Number of Tickets")
		fmt.Scan(&ticketNumber)

		if ticketNumber <= availableTicket {
			availableTicket = availableTicket - ticketNumber
			// bookings[0]= firstName+" "+lastName

			bookings = append(bookings, firstName+" "+lastName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("Thank you %v %v for purchasing %v ticket,your confirmation mail will be sent to the email address you provided %v.\n", firstName, lastName, ticketNumber, email)
			fmt.Printf("%v remaining tickets for the %v.\n", availableTicket, conferenceName)
			fmt.Printf("This are the first names for our bookings: %v\n", firstNames)

			if availableTicket == 0 {
				fmt.Println("The Conference ticket are booked out,come back next year")
				break
			}

		} else {
			fmt.Printf("We only have %v remaining, so kindly reajust your ticket No \n", availableTicket)
		}

	}

	//fmt.Println(bookings)
	//fmt.Println(len(bookings))
	//fmt.Println(bookings[0])

}
