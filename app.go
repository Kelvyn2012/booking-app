package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go-conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	// fmt.Println("Welcome to", conferenceName, "booking application!!")
	// fmt.Println("Get your tickets here to attend")
	//fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", conferenceTickets, remainingTickets)
	greetUsers(conferenceName,conferenceTickets,remainingTickets)

	for {
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

		// input validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

		// using slice instead of an array
		

		if isValidUserTickets && isValidEmail && isValidName {
			bookings = append(bookings, firstName+" "+lastName)
			remainingTickets = remainingTickets - userTickets

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("This are the first names for all our bookings %v \n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference Tickets is booked out, kindly come back nect year!!")
				break
			}

		} else {
			if !isValidEmail{
				fmt.Println("Your email ID does not contain @")
			}
			if !isValidName{
				fmt.Println("Your first name or last name is too short")
			}
			if !isValidUserTickets{
				fmt.Println("Your ticket number is either too low and above number of ticket available")
			}
			continue
		}
	}

}

func greetUsers(confName string,confTickets uint,remainTickets uint){
	fmt.Printf("Welcome to %v booking application!!\n", confName)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", confTickets, remainTickets)
}