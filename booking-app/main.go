package main

import "fmt"

func main() {
	const conferenceTickets = 50

	var (
		conferenceName = "Go Conference"
		remainingTickets uint = conferenceTickets
		bookings []string

		firstName string
		lastName string
		email string
		ticketCount uint 
	)

	fmt.Printf("Type of remiainingTickets is %T\n", remainingTickets)


	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")


	//User Inputs

	fmt.Printf("Enter you first name : ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter you last name : ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter you email address : ")
	fmt.Scan(&email)

	fmt.Printf("Enter number of tickets yuo want to buy : ")
	fmt.Scan(&ticketCount)

	// Updates
	remainingTickets = remainingTickets - ticketCount
	bookings = append(bookings, firstName + " " + lastName)


	// Results
	
	fmt.Printf("Thank yuu %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, ticketCount, email)
	fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, conferenceName)
    fmt.Printf("These are all out bookings : %v\n", bookings)	
}
