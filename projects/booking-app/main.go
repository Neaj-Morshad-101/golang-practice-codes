package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = 50
	bookings              = make([]UserData, 0)
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		userData := getUserInput()
		isValid := validateUserInput(userData)

		if isValid {
			bookTicket(userData)
			wg.Add(1) // //increment waiting go routing count

			go sendTicket(userData)

			printFirstNames()

			// TODO: CHECK FROM HERE

			if remainingTickets == 0 {
				//end program
				fmt.Printf("Our %v is booked out. Come back next year.\n\n\n", conferenceName)
				break
			}
		}
	}
	//wait for go routine to finish before terminating the program (the main function).
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getUserInput() UserData {
	var userData UserData

	fmt.Printf("Enter you first name : ")
	fmt.Scan(&userData.firstName)

	fmt.Printf("Enter you last name : ")
	fmt.Scan(&userData.lastName)

	fmt.Printf("Enter you email address : ")
	fmt.Scan(&userData.email)

	fmt.Printf("Enter number of tickets yuo want to buy : ")
	fmt.Scan(&userData.numberOfTickets)

	return userData
}

func bookTicket(userData UserData) {
	remainingTickets = remainingTickets - userData.numberOfTickets
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n\n", userData.firstName, userData.lastName, userData.numberOfTickets, userData.email)
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("%v tickets remaining for %v\n\n\n", remainingTickets, conferenceName)

}

func sendTicket(userData UserData) {
	time.Sleep(50 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userData.numberOfTickets, userData.firstName, userData.lastName)
	fmt.Println("----------------------")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, userData.email)
	fmt.Println("----------------------")
	wg.Done() //decrement waiting go routing count
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}
