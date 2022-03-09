package main

import (
	"fmt"
	"strconv"
)

var bookingName = "Rub Booking"

const numTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidNum := validateUserInput(firstName, lastName, email, userTickets)

		if !isValidNum {
			fmt.Printf("There are only %v tickets remaining. %v is greater than %v\n", remainingTickets, userTickets, remainingTickets)
			continue
		}
		if !isValidName {
			fmt.Println("First or last name is not valid. Please try again")
			continue
		}
		if !isValidEmail {
			fmt.Println("Email is invalid. Please try again")
			continue
		}

		bookTickets(userTickets, firstName, lastName, email)
		firstNames := printFirstNames()
		fmt.Printf("Total bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("There are no more tickets left. Thank you for booking with us!")
			break
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", bookingName)
	fmt.Printf("We have %v tickets left out of %v total tickets\n", remainingTickets, numTickets)
	fmt.Println("Book your tickets here")
	fmt.Println(bookingName)
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = booking["firstName"]
		firstNames = append(firstNames, name)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Please enter your email: ")
	fmt.Scan(&email)
	fmt.Printf("Hello %v %v, please enter how many tickets you would like: ", firstName, lastName)
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// Create a user map
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings:\n %v", bookings)

	fmt.Printf("Ordering %v tickets for %v %v. Order confirmation sent to %v\n", userTickets, firstName, lastName, email)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
}
