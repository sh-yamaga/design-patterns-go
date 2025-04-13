# Memento Pattern

## What is the "Memento Pattern"?

The Memento Pattern is a behavioral design pattern that provides the ability to capture and restore an object's internal state without exposing its implementation details. By storing snapshots (mementos) of an object's state, you can easily revert or roll back to a previous state later.

![Class Diagram](./assets/class-diagram.drawio.png)

## Why is the "Memento Pattern" needed?

1. **State Preservation and Undo Functionality**
It allows you to implement features like “Undo” or “Rollback” by capturing the state at different points in time.

2. **Encapsulation of Internal State**
The Originator (the object whose state you want to save) does not have to expose its internal fields, preventing outside code from relying on fragile implementation details.

3. **Ease of Maintenance**
Each state snapshot is a self-contained object. These snapshots can be created or destroyed without affecting unrelated code, simplifying your overall design.

4. **Reduced Complexity**
By moving the responsibility for saving and restoring state into separate classes (the Memento and the Caretaker/History), you avoid cluttering the main object's implementation with undo/redo logic.

## Sample Program

Below is a simplified example of the Memento Pattern in action, featuring a text editor capable of saving and restoring its content:

- **Editor**(Originator): Maintains the current text and cursor position. It can create a memento (Save()) capturing its current internal state and restore a previous state (Restore()).

- **EditorMemento**(Memento): Holds a snapshot of the Editor’s internal state (the text and the cursor position), which is used to restore that state later.

- **History**(Caretaker): Maintains a list (stack) of mementos and provides ways to save (push) or restore (pop) previous states of the editor.

To run the sample code, execute 
```bash
go run ./cmd/memento/main.go
```