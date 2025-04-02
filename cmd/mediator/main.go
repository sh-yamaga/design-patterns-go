package main

func main() {
	chatRoom := mediator.ChatRoom{}

	userA := &mediator.User{name: "Harry", madiator: chatRoom}
	userB := &mediator.User{name: "Ron", madiator: chatRoom}
	userC := &mediator.User{name: "Hermione", madiator: chatRoom}

	chatRoom.AddUser(userA)
	chatRoom.AddUser(userB)
	chatRoom.AddUser(userC)

	userC.SendMessage("Holy cricket. You're Harry Potter!")
	userC.SendMessage("I'm Hermione Granger. And... you are?")
	userB.SendMessage("Um... Ron Weasley.")
	userC.SendMessage("Preasure.")
}
