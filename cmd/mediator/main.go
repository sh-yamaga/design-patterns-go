package main

import "github.com/sh-yamaga/design-patterns-go/internal/patterns/mediator"

func main() {
	chatRoom := &mediator.ChatRoom{}

	userA := mediator.NewUser("Harry", true, chatRoom)
	userB := mediator.NewUser("Ron", true, chatRoom)
	userC := mediator.NewUser("Hermione", false, chatRoom)

	chatRoom.AddUser(userA)
	chatRoom.AddUser(userB)
	chatRoom.AddUser(userC)

	userC.SendMessage("Holy cricket. You're Harry Potter!")
	// Output:
	// Hermione send a message: Holy cricket. You're Harry Potter!
	// Harry recieved a message from Hermione: Holy cricket. You're Harry Potter!
	// Ron recieved a message from Hermione: Holy cricket. You're Harry Potter!

	userC.SendMessage("I'm Hermione Granger. And... you are?")

	userB.SendMessage("Um... Ron Weasley.")
	// Output:
	// Ron send a message: Um... Ron Weasley.
	// Harry recieved a message from Ron: Um... Ron Weasley.

	userC.SendMessage("Preasure.")
}
