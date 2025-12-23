package main

import (
	"fmt"
	"strings"
)



func main(){
	var conferenceName string = "Go-conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Println("Welcome to", conferenceName ,"booking application!!")
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", conferenceTickets, remainingTickets)

	

	for{
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your Email")
		fmt.Scan(&email)
		fmt.Println("Number of Tickets")
		fmt.Scan(&userTickets)

	// using slice instead of an array
		bookings = append(bookings, firstName + " " + lastName)

		if userTickets <= remainingTickets{
			remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		
		firstNames := []string{}

		for _,booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("This are the first names for all our bookings %v \n",firstNames)
		if remainingTickets == 0{
			fmt.Println("Our conference Tickets is booked out, kindly come back nect year!!")
			break
		}


	}else{
		fmt.Printf("We only have %v remaining tickets,so you cant book %v tickets \n",remainingTickets,userTickets)
		continue
	}
		}

		

		

}