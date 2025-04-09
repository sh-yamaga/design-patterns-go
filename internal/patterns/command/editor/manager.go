package editor

import "fmt"

type Command interface {
	Execute()
	Undo()
}

// CmdManager manages excecution and undo of commands
type CmdManager struct {
	history []Command
}

// ExcecuteCommand excecutes command and record its history
func (cm *CmdManager) ExecuteCommand(cmd Command) {
	cmd.Execute()
	cm.history = append(cm.history, cmd)
}

func (cm *CmdManager) Undo() {
	if len(cm.history) == 0 {
		fmt.Println("There is no command to undo")
		return
	}

	lastCmd := cm.history[len(cm.history)-1]
	lastCmd.Undo()

	// remove lastCmd from command history
	cm.history = cm.history[:len(cm.history)-1]
}
