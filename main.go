package main

import (
	"booking-app/helper"
	"fmt"
)
	var conferenceName string = "Go-conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings = make([]userData,0)


	type userData struct{
		firstName string
		lastName string
		Email string
		numberOfTickets uint
	}

func main() {
	

	greetUsers()

	for {
		// user details
		firstName,lastName,email,userTickets := getUserDetails()

		// input validation
		isValidUserTickets,isValidEmail,isValidName := helper.ValidateUserInput(userTickets,firstName,lastName, email,remainingTickets)

		// using slice instead of an array
		

		if isValidUserTickets && isValidEmail && isValidName {
			bookTickets(userTickets,firstName,lastName,email)

			// print first names
			firstNames := getFirstName()
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

func greetUsers(){
	fmt.Printf("Welcome to %v booking application!!\n", conferenceName)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("We have total of %v Tickets to be sold, and we have %v Tickets available!!\n", conferenceTickets, remainingTickets)
}
func getFirstName()[]string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
		}
		return firstNames
		
}
func getUserDetails()(string,string,string,uint){
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
		return firstName,lastName,email,userTickets
}
func bookTickets(userTickets uint,firstName string, lastName string, email string){
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		Email: email,
		numberOfTickets: userTickets,
		
		
	}

	bookings = append(bookings, userData)
	fmt.Printf("Details of bookings are %v\n",bookings)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}







