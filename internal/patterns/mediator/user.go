package mediator

import "fmt"

type User struct {
	name         string
	notification bool
	mediator     *ChatRoom
}

func NewUser(name string, notification bool, mediator *ChatRoom) *User {
	return &User{
		name:         name,
		notification: notification,
		mediator:     mediator,
	}
}

func (u *User) SendMessage(msg string) {
	fmt.Printf("%s send a message: %s\n", u.name, msg)
	u.mediator.SendMessage(msg, u)
}

func (u *User) RecieveMessage(msg string, sender *User) {
	fmt.Printf("%s recieved a message from %s: %s\n", u.name, sender.name, msg)
}
