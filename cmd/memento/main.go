package main

import (
	"fmt"

	"github.com/yamaga-shu/design-patterns-go/internal/patterns/memento"
)

func main() {
	editor := memento.NewEditor()
	history := memento.History{}

	// initial state
	editor.TypeWords("Hello")
	history.Push(editor.Save())

	// 2nd state
	editor.TypeWords(", World! ")
	history.Push(editor.Save())

	// 3rd state
	editor.TypeWords("This is a Memento Pattern Sample.")
	history.Push(editor.Save())

	// print current(3rd) state
	fmt.Printf("3rd Text	: \"%s\"\n", editor.Text)
	fmt.Printf("3rd Posi	: %d\n", editor.CursorPosition)
	// Output:
	// 3rd Text	: "Hello, World! This is a Memento Pattern Sample."
	// 3rd Posi	: 47

	// type wrong word
	editor.TypeWords("wrong text")

	// restore 3rd state
	/// poping
	lastState := history.Pop()
	/// restoring
	if lastState != nil {
		editor.Restore(lastState)
	}
	/// print 3nd state again
	fmt.Printf("3nd Text	: \"%s\"\n", editor.Text)
	fmt.Printf("3nd Posi	: %d\n", editor.CursorPosition)
	// Output:
	// 3rd Text	: "Hello, World! This is a Memento Pattern Sample."
	// 3rd Posi	: 47

	// restore 2nd state
	/// poping
	lastState = history.Pop()
	/// restoring
	if lastState != nil {
		editor.Restore(lastState)
	}
	/// print 2nd state
	fmt.Printf("2nd Test	: \"%s\"\n", editor.Text)
	fmt.Printf("2nd Posi	: %d\n", editor.CursorPosition)
	// Output:
	// 2nd Test	: "Hello, World! "
	// 2nd Posi	: 14
}
