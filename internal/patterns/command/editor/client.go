package editor

import "fmt"

type Command interface {
	Execute()
	Undo()
}

// EditorClient manages excecution and undo of commands
type EditorClient struct {
	history []Command
}

// ExcecuteCommand excecutes command and record its history
func (ec *EditorClient) ExecuteCommand(cmd Command) {
	cmd.Execute()
	ec.history = append(ec.history, cmd)
}

func (ec *EditorClient) Undo() {
	if len(ec.history) == 0 {
		fmt.Println("There is no command to undo")
		return
	}

	lastCmd := ec.history[len(ec.history)-1]
	lastCmd.Undo()

	// remove lastCmd from command history
	ec.history = ec.history[:len(ec.history)-1]
}
