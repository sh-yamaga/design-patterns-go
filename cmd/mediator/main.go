package main

func main() {
	chatRoom := mediator.ChatRoom{}

	userA := &mediator.User{
		name:         "Harry",
		notification: true,
		madiator:     chatRoom,
	}
	userB := &mediator.User{
		name:         "Ron",
		notification: true,
		madiator:     chatRoom,
	}
	userC := &mediator.User{
		name:         "Hermione",
		notification: false,
		madiator:     chatRoom,
	}

	chatRoom.AddUser(userA)
	chatRoom.AddUser(userB)
	chatRoom.AddUser(userC)

	userC.SendMessage("Holy cricket. You're Harry Potter!")
	userC.SendMessage("I'm Hermione Granger. And... you are?")
	userB.SendMessage("Um... Ron Weasley.")
	userC.SendMessage("Preasure.")
}
