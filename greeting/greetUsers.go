package greeting

import (
	"fmt"
)

const conferenceTickets uint = 250

func GreetUsers(comedian, conferenceName string, remainingTickets, count uint) {
	fmt.Printf("%v welcomes you to the show\n", comedian)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
