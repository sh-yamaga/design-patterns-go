package mediator

type ChatRoom struct {
	user []*User
}

func (cr *ChatRoom) AddUser(u *User) {
	cr.user = append(cr.user, u)
}

func (cr *ChatRoom) SendMessage(msg string, sender *User) {
	for _, user := range cr.user {
		if user != sender && user.notification {
			user.RecieveMessage(msg, sender)
		}
	}
}
