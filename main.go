package main

import "fmt"

func main() {
	var conferenceName = "Go-conference"
	const ticketNumber = 50
	var availableTicket = 50

	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", ticketNumber, availableTicket)
	println("Get your tickets here to attend")

}
