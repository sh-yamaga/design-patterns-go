# Decorator Pattern

## What is the "Decorator Pattern"?

The Decorator Pattern is a structural design pattern that allows you to attach new behaviors to objects by placing these objects inside special wrapper objects. Decorators provide a flexible alternative to subclassing for extending functionality. Rather than creating multiple subclasses, you can add additional behaviors dynamically by nesting decorators.

![Class Diagram](./assets/class-diagram.drawio.png)

## Why is the "Decorator Pattern" needed?

1. **Flexible Extension of Functionality**
You can add new capabilities to objects without changing their underlying code or creating an entire hierarchy of subclasses.

2. **Open/Closed Principle**
The Decorator Pattern allows a class to remain closed for modification (by avoiding changes to the original class), while open for extension (by creating decorators that implement the same interface).

3. **Simplified Maintenance**
Instead of creating multiple subclasses to handle each new feature, you can create decorators that wrap a base object, making it easier to maintain and extend.

4. **Layered Enhancements**
Decorations can be composed by nesting them, stacking multiple behaviors in a single pipeline.

## Sample Program

Below is a simplified example of how you can use the Decorator Pattern for stream processing:

- **Interface**: The common interface (Interface) that all components (including the decorators) implement.

- **Stream`** (Concrete Component): Contains the original data.
  
- **UpperCaseFilter** (Concrete Decorator): Wraps a Stream and transforms the data to uppercase before returning it.

By organizing handlers in a chain, you can easily add, remove, or shift the responsibility of different HTTP status codes without changing the client code that sends these responses.

To run the sample code, execute 
```bash
go run ./cmd/decorator/main.go
```