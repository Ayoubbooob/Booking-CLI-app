package user

import "strings"

func ValidateUserInputs(userFirstName string, userLastName string, email string, tickets uint, ticketAvailable uint) (bool, bool, bool) {
	isValidName := len(userFirstName) > 2 && len(userLastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsRequest := tickets <= ticketAvailable

	return isValidName, isValidEmail, isValidTicketsRequest
}
