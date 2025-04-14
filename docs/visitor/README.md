# Visitor Pattern

## What is "Visitor Pattern" ?

The Visitor Pattern is a behavioral design pattern that lets you separate algorithms from the objects on which they operate. In other words, it allows adding new operations to existing object structures without modifying those structures.

## Why is "Visitor Pattern" needed ?

1. **Separation of Concerns**: By keeping the operations separate from the data structures, each type of operation (visitor) can evolve independently.

2. **Extensibility**: You can define new visitors (operations) without modifying existing object classes.

3. **Single Responsibility**: Each visitor focuses on a specific task or operation, while the elements themselves remain focused on their primary responsibilities.

4. **Maintainability**: Related operations are encapsulated in separate visitors, making them easier to locate, understand, and change.

## Sample Program

This sample code demonstrates the use of the Visitor Pattern with geometric shapes. Two visitors are defined:
- **DisplayVisitor**: Displays each shape’s characteristics (e.g., radius, width, height).  
- **AreaVisitor**: Calculates and prints the area of the shapes.

- **Elements**: In this example, `Circle` and `Rectangle` implement an `Element` interface and accept visitors.  
- **Visitor Interface**: Declares visit methods for different elements such as `VisitCircle` and `VisitRectangle`.  
- **Concrete Visitors**: `DisplayVisitor` and `AreaVisitor` implement the interface methods for each shape.

To run the sample code, execute:
```bash
go run ./cmd/visitor/main.go
```
