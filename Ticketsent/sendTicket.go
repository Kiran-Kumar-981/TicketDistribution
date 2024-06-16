package Ticketsent

import (
	"fmt"
	"time"
)

func SendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(2 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	fmt.Println(
		`your show will be held on Monday(dd-mm-yyyy)
		 please come 15-min early for ticket checking`)
}
