package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/command"
	"github.com/yamaga-shu/design-patterns-go/internal/patterns/command/editor"
)

func main() {
	e := editor.NewEditor()
	cmdManager := &editor.CommnadManager{}

	cmdManager.ExecuteCommand(&command.InsertCommand{
		Editor:       e,
		CursorIndex:  0,
		TextToInsert: "Hello, ",
	})
	cmdManager.ExecuteCommand(&command.InsertCommand{
		Editor:       e,
		CursorIndex:  len(e.Text()),
		TextToInsert: "World!",
	})

	fmt.Println(e.Text())
	// Output:
	// Hello, World!

	// delete "!"
	cmdManager.ExecuteCommand(&command.DeleteCommand{
		Editor:      e,
		CursorIndex: len(e.Text()) - 1,
		Length:      1,
	})

	fmt.Println(e.Text())
	// Output:
	// Hello, World

	// Undo: Cancel the last DeleteCommand
	cmdManager.Undo()
	fmt.Println("Canceled DeleteCommand:", e.Text())
	// Output:
	// Canceled DeleteCommand: Hello, World!

	// Undo: Cancal the last InsertCommand
	cmdManager.Undo()
	fmt.Println("Canceled last InsertCommand:", e.Text())
	// Output:
	// Canceled last InsertCommand: Hello,

	// Undo: Cancel the first InsertCommand
	cmdManager.Undo()
	fmt.Println("Canceled first InsertCommand:", e.Text())
	// Output:
	// Canceled first InsertCommand:

	// Undo: Noting happened
	cmdManager.Undo()
	// Output:
	// There is no command to undo
}
