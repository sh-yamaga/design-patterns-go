# Facade Pattern

## What is the "Facade Pattern"?

The Facade Pattern is a structural design pattern that provides a simplified interface to a complex subsystem. By introducing a facade, you can unify multiple subsystems or components under a single interface, making it easier for external clients to use the underlying functionality without needing to understand the system’s details.

![Class Diagram](./assets/class-diagram.drawio.png)

## Why is the "Facade Pattern" needed?

1. **Simplified Interface**
The facade provides a single point of entry to a set of classes or modules, allowing clients to interact with the subsystem through concise methods instead of multiple dependencies.

2. **Decoupling Subsystems**
By interacting only with the facade, client code remains decoupled from the specific classes and their interactions within the subsystem, improving maintainability and scalability.

3. **Easier to Extend & Maintain**
Adding or modifying subsystem behavior can often be handled by adjusting the facade without requiring sweeping changes in client code.

4. **Improved Readabilitys**
Clients deal with straightforward methods that describe higher-level operations rather than low-level implementation details.

## Sample Program

In this project, the facade TravelFacade encapsulates the steps needed to book a flight, reserve a hotel, and arrange a rental car through individual booker modules. Rather than invoking each booker yourself, you can simply call one method on the facade:

- **TravelFacade**(Facade class): Wraps multiple booking services (FlightBooker, HotelBooker, RentalCarBooker) and exposes a single high-level interface (BookTravel).

- **FlightBooker, HotelBooker, RentalCarBooker**: Subsystem classes responsible for their respective booking tasks.

- **Client Code**(in main.go): Creates a TravelFacade instance and calls BookTravel("Tokyo") once to handle all the bookings.

To run the sample code, execute 
```bash
go run ./cmd/facade/main.go
```