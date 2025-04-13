# Mediator Pattern

## What is the "Mediator Pattern"?

The Mediator Pattern is a behavioral design pattern that defines an object (the mediator) that encapsulates how a set of objects interact. It promotes loose coupling by keeping objects from referring to each other explicitly, enabling you to vary their interactions independently.

![Class Diagram](./assets/class-diagram.drawio.png)

## Why is the "Mediator Pattern" needed?

1. **Reduced Coupling**
Instead of components calling each other directly, they communicate through a centralized mediator. Each component only needs to know about the mediator, not the other components.

2. **Simplified Communication**
The mediator provides a single point through which messages can be passed or operations orchestrated, making interactions easier to follow and debug.

3. **Flexibility and Maintainability**
You can add new components or change existing ones without breaking the entire system’s communication logic. The central mediator manages inter-component relationships.

4. **Easier Collaboration**
With a clear mediator, you can focus on each component’s functionality rather than worrying about keeping track of many cross-object references.

## Sample Program

In this repository, an example of the Mediator Pattern is demonstrated by a simple chat application:

- **ChatRoom**(Mediator): Holds a list of users in the chat. It implements the logic for broadcasting messages to all users except the sender.

- **User**(Colleague): Each user has a reference to the mediator (ChatRoom) and can send or receive messages. When a user sends a message, it calls the mediator’s method to distribute that message to other users.

- **AndExpression & OrExpression**(Composite Expressions): Evaluate multiple expressions in combination.
  - `AndExpression` returns true only if both sub-expressions are true.
  - `OrExpression` returns true if at least one of the sub-expressions is true.

This approach cleanly decouples each User object from every other User, since communication details remain in the ChatRoom mediator.

To run the sample code, execute 
```bash
go run ./cmd/mediator/main.go
```