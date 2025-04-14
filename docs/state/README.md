# State Pattern

## What is "State Pattern" ?

The State Pattern is a behavioral design pattern that allows an object to alter its behavior when its internal state changes. The object essentially appears to change its class. This helps in organizing code around distinct states, making logic easier to maintain.

## Why is "State Pattern" needed ?

1. **Controlled Transitions**: Rather than large if-else or switch statements, each state encapsulates its own behavior and defines valid transitions to other states.  

2. **Maintainability**: Changes to one state’s logic don’t affect other states, preventing complex conditional branches from becoming unmanageably large.  

3. **Separation of Responsibilities**: States are self-contained, making it clearer how the system behaves in each situation.  

4. **Flexibility**: Adding or modifying states can be done with fewer changes to the rest of the codebase.

## Sample Program

This sample code demonstrates the use of the State Pattern with a vending machine, which has the following states:
- **NoCoinState**: Ready to accept a coin.  
- **HasCoinState**: A coin has been inserted; ready to select a product.  
- **DispendingState**: Currently dispensing the product.  
- **SoldOutState**: No products left.

The `VendingMachine` transitions between states to manage product dispensing.

To run the sample code, execute:
```bash
go run ./cmd/state/main.go
```
