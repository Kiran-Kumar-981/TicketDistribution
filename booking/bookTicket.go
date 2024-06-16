package booking

import (
	"fmt"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var remainingTickets uint
var bookings = make([]UserData, 0)

func BookTicket(userTickets uint, firstName, lastName, email, conferenceName string) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf(`Thank you %v %v for booking %v tickets. 
	You will receive a confirmation email at %v`, firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
func GetFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
