# Prototype Pattern

## What is "Prototype Pattern" ?

The Prototype Pattern is a creational design pattern that enables cloning of existing objects without making the code dependent on their classes. By creating copies from prototypes, you avoid the overhead of creating objects from scratch, including complex initialization, while still controlling which objects are copied.

## Why is "Prototype Pattern" needed ?

1. **Efficiency**: Cloning an existing object can be more cost-effective than creating a new one from scratch (especially when initialization is time-consuming or relies on complex data).

2. **Simplified Object Creation**: When object creation involves numerous parameters or extensive configuration, using a prototype can simplify the process by encapsulating that complexity in the object itself.

3. **Decoupling**: The pattern decouples object creation from its specific class, allowing new prototypes or subtypes to be added without modifying existing code.

4. **Run-time Flexibility**: You can change and reconfigure prototypes at run time, adding or removing prototypes as needed without impacting other parts of the application.

## Sample program

This sample code demonstrates the use of the Prototype Pattern with geometric shapes (a triangle and a square). You can create clones of shapes without dealing with their specific implementations.

- **Shape (Interface)**: Defines the contract for both drawing and cloning.  
- **Triangle & Square**: Concrete prototypes that implement the shape interface. They each define how they draw themselves and how they create new instances of the same type.

To run the sample code, execute:
```bash
go run ./cmd/prototype/main.go
```