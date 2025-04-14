# Singleton Pattern

## What is "Singleton Pattern" ?

The Singleton Pattern is a creational design pattern that ensures only one instance of a particular class is created during the lifetime of an application. It provides a global point of access to that instance, while preventing multiple instances from being created.

## Why is "Singleton Pattern" needed ?

1. **Controlled Access**: Ensures that critical resources are managed by a single instance, preventing conflicts or inconsistencies.  

2. **Resource Efficiency**: Avoids the overhead of creating multiple instances of objects that share the same data or state.  

3. **Simplicity**: Provides a straightforward and well-defined point of access when exactly one instance is required.  

4. **Centralized Management**: Allows global configuration or state to be maintained in one place.

## Sample Program

This sample code demonstrates the use of the Singleton Pattern with a ticket manager. You can generate ticket numbers from a single source, ensuring they are unique and sequential.

- **TicketManager**: The global resource that provides sequential ticket numbers.  
- **NewTicketManager**: Returns a single shared instance of `TicketManager` using the lazy initialization and thread-safe approach with the `sync.Once` pattern.  
- **Next**: Increments and returns the next ticket number in sequence.

To run the sample code, execute:
```bash
go run ./cmd/singleton/main.go
```