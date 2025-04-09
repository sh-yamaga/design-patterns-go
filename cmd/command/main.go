package main

import (
	"fmt"

	"github.com/sh-yamaga/design-patterns-go/internal/patterns/command"
	"github.com/sh-yamaga/design-patterns-go/internal/patterns/command/editor"
)

func main() {
	e := editor.NewEditor()
	cmdManager := &editor.CmdManager{}

	cmdManager.ExecuteCommand(&command.InsertCommand{
		Editor:         e,
		CursorPosition: 0,
		TextToInsert:   "Hello, ",
	})

	cmdManager.ExecuteCommand(&command.InsertCommand{
		Editor:         e,
		CursorPosition: len(e.Text()),
		TextToInsert:   "World!",
	})

	fmt.Println(e.Text())
	// Output:
	// Hello, World!
}
