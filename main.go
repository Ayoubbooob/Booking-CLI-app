package main

import (
	"booking-app/user"
	"fmt"
	"strings"
	"time"
)

const conferenceName = "GoLang Geeks"

var ticketAvailable uint = 50

var userID int = -1

var bookingUsers = []UserData{}

type UserData struct {
	firstname           string
	lastname            string
	email               string
	numberOfUserTickets uint
}

func main() {

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	for next := true; next; next = ticketAvailable > 0 {

		greetingUser()

		userFirstName, userLastName, email, tickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketsRequest := user.ValidateUserInputs(userFirstName, userLastName, email, tickets, ticketAvailable)

		if isValidEmail && isValidName && isValidTicketsRequest {

			userID++

			userData := UserData{
				firstname:           userFirstName,
				lastname:            userLastName,
				email:               email,
				numberOfUserTickets: tickets,
			}

			bookingUsers = append(bookingUsers, userData)

			fmt.Println("")
			fmt.Println("")

			fmt.Println("")
			fmt.Println("")
			fmt.Println("")
			fmt.Println("")

			ticketAvailable -= tickets

			go sendTicket(userData.firstname, userData.lastname, userData.email, userData.numberOfUserTickets)
			fmt.Printf("							 🎉 🎉 🎉 Yuppi, Congratulation %v %v. You will be with us ! 🎊 🎊 🎊", bookingUsers[userID].firstname, bookingUsers[userID].lastname)

			fmt.Println("")
			fmt.Println("")
			if tickets == 1 {
				fmt.Println("							 🎉 🎉 🎉 You will receive your ticket in your email", email, "🎊 🎊 🎊")
			} else {
				fmt.Println("							 🎉 🎉 🎉 You will receive your", tickets, "tickets in the email", email, " 🎊 🎊 🎊")
			}

			fmt.Println("")
			fmt.Println("")
			fmt.Println("")
			fmt.Println("")
			fmt.Println("")
			fmt.Println("")

		} else {

			if !isValidEmail {
				fmt.Println("							❌ ❌ ❌ Your email is wrong, Try again please ! ❌ ❌ ❌", ticketAvailable, tickets)
				fmt.Print("Your Email : ")
				fmt.Scan(&email)
				fmt.Println("")
				fmt.Println("")
			}

			if !isValidName {
				fmt.Println("							❌ ❌ ❌ Firstname & Lastname should be at least of 3 characters ! ❌ ❌ ❌")
				fmt.Print("											Your Firstname : ")
				fmt.Scan(&userFirstName)
				fmt.Println("")
				fmt.Print("											Your Lastname : ")
				fmt.Scan(&userLastName)
				fmt.Println("")
				fmt.Println("")

			}

			if !isValidTicketsRequest {
				fmt.Printf("							❌ ❌ ❌ You've requested %v tickets, but just %v are remaining ! ❌ ❌ ❌", ticketAvailable, tickets)
				fmt.Print("											Your Email : ")
				fmt.Scan(&email)
				fmt.Println("")
				fmt.Println("")
			}
		}
		fmt.Println("  											👥 👥 👥	The First names of bookings are  : ", getBookingNames())
		fmt.Println("")
		fmt.Println("")

	}

}

func greetingUser() {
	fmt.Println(" 									😄 😄 😄 Welcome to ", conferenceName, "booking app ! 😄 😄 😄	")
	fmt.Println("")
	fmt.Println("")
}

func getUserInputs() (string, string, string, uint) {

	var userFirstName string
	var userLastName string
	var email string
	var tickets uint

	fmt.Println("				 **** let us know you more !  	")

	fmt.Println("")
	fmt.Println("")

	fmt.Print("						* What's your firstname : ")
	fmt.Scan(&userFirstName)
	fmt.Print(" 						* What's your lastname : ")

	fmt.Scan(&userLastName)

	fmt.Println("")
	fmt.Println("")

	fmt.Print(" 						* What's your email : ")
	fmt.Scan(&email)

	fmt.Println("")
	fmt.Println("")

	fmt.Println(" 								⌛ ⌛ ⌛ Tickets Available :", ticketAvailable, "Harry up before we close ! ⌛ ⌛ ⌛")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("								✅ ✅	✅ How many tickets you want get ? ✅ ✅ ✅")
	fmt.Print("														")

	fmt.Scan(&tickets)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	return userFirstName, userLastName, email, tickets
}

func getBookingNames() []string {

	bookingFirstNames := []string{}

	for _, booking := range bookingUsers {
		var names = strings.Fields(booking.firstname)
		bookingFirstNames = append(bookingFirstNames, names[0])
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	return bookingFirstNames
}

func sendTicket(firstname string, lastname string, email string, tickets uint) {

	time.Sleep(5 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", tickets, firstname, lastname)
	fmt.Println("								########################################################################## ")
	fmt.Println("")
	fmt.Println("")

	fmt.Printf("									Sending tickets: ...\n")
	fmt.Printf("										%v to email address %v \n", ticket, email)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("								##########################################################################")

	fmt.Println("")
	fmt.Println("")

}
