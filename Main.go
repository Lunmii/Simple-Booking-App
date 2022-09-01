package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	age             uint
	email           string
	numberofTickets uint
}

func main() {

	greetUsers()

	for {
		firstName, lastName, age, email, userTickets := getUserInput()
		isValidName, isValidAge, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, age, email, userTickets, remainingTickets)

		if isValidName && isValidAge && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, age, email, userTickets)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName()
			fmt.Printf("The First names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is to short")
			}

			if !isValidAge {
				fmt.Printf("According to the age you entered, you're not allowed to attend %v\n", conferenceName)
			}

			if !isValidEmail {
				fmt.Println("Email Address is incorrect, Try Again!")
			}

			if !isValidTicketNumber {
				fmt.Printf("Insufficient number of tickets, Only %v tickets are available\n", remainingTickets)
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, bookings := range bookings {
		firstNames = append(firstNames, bookings.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, uint, string, uint) {

	var firstName string
	var lastName string
	var age uint
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, age, email, userTickets
}

func bookTicket(firstName string, lastName string, age uint, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		age:             age,
		email:           email,
		numberofTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation message at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###########")
}
