# Observer Pattern

## What is "Observer Pattern" ?

The Observer Pattern is a behavioral design pattern that establishes a one-to-many relationship between objects. When one object (the “subject”) changes its state, all its dependents (the “observers”) are automatically notified and updated. This promotes a loose coupling between the subject and its observers.

![Class Diagram](./assets/class-diagram.drawio.png)

## Why is "Observer Pattern" needed ?

1. **Loose Coupling**: Observers aren’t tightly coupled to the subject, allowing either party to evolve independently without breaking the other.

2. **Automatic Updates**: When the subject changes, it notifies all observers, ensuring that dependent data stays consistent without additional overhead in each observer.

3. **Flexibility**: New observers can be added or removed at runtime without modifying the subject, making it easier to extend or change program behavior.

4. **Reusability**: By abstracting communication through a common interface, the same subject can be reused with different observer implementations.

## Sample program

![Sample program diagram](./assets/sample-program.drawio.png)

This sample code demonstrates the use of the Observer Pattern with a weather station example. The program allows you to register multiple observers (like a current conditions display and a statistics display) that are notified whenever new weather measurements are set on the subject.

- **Weather (Subject)**: Maintains current weather data and notifies registered observers about state changes.  
- **Observer Interface**: A contract that all observers implement to receive updated weather values.  
- **CurrentDisplay & StatisticsDisplay**: Concrete observers that implement the observer interface. They each maintain and display weather data in different ways.

To run the sample code, execute:
```bash
go run ./cmd/observer/main.go
```