package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings []string
// var bookings = []string{}
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//user input validation
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	//user have to choose location
	// isValidCity := city == "London" || city == "Munich"

	//above three have to true

	if isValidName && isValidEmail && isValidTicketsNumber {

		bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		//call function print first names
		firstNames := getFirstNames()
		fmt.Printf("These first names of bookings are: %v \n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			//end program
			fmt.Println("Our conference is booked out come next year")

		}
	} else {

		//generic message which shows input is invalid
		if !isValidName {
			fmt.Println("first name or last name you enterd is too short.")
		}
		if !isValidEmail {
			fmt.Println("The sequence of email is not right.")
		}
		if !isValidTicketsNumber {
			fmt.Println("number of tickets you enterd is invalid")
		}
	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask used for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets,. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######")
	fmt.Printf("Sending ticket: \n %v to email address %v\n", ticket, email)
	fmt.Println("######")
	wg.Done()
}
