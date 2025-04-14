# Strategy Pattern

## What is "Strategy Pattern" ?

The Strategy Pattern is a behavioral design pattern that defines a family of algorithms, encapsulates each one, and makes them interchangeable at run time. This allows the algorithm to vary independently from the clients that use it, promoting cleaner code organization and easier code maintenance.

## Why is "Strategy Pattern" needed ?

1. **Flexibility**: By encapsulating specific algorithms, you can change the strategy without modifying the client code.

2. **Maintainability**: Each algorithm is defined in its own type, simplifying updates and organization.

3. **Extensibility**: You can add new strategies or remove old ones without altering existing client code.

4. **Reusability**: Strategies can be reused across different contexts by adhering to a common interface.

## Sample Program

This sample code demonstrates the use of the Strategy Pattern with a guessing game example. The program randomly generates numbers, then uses two strategies to guess those numbers.

- **BruteForce**: Implements a brute-force approach, checking each incrementally higher value until it finds the target.  
- **BinarySearch**: Implements a more efficient binary search algorithm.  
- **Strategy Interface**: Provides a contract (`GuessNumber` & `String`) that both strategies implement.  

The main program measures how long it takes for each strategy to guess a large number of random values.

To run the sample code, execute:
```bash
go run ./cmd/strategy/main.go
```
