# Interpreter Pattern

## What is the "Interpreter Pattern"?

The Interpreter Pattern is a behavioral design pattern that defines a representation for a language’s grammar along with an interpreter to process that grammar. It allows you to interpret sentences (or expressions) within a simple language by defining classes for each rule in the grammar. The pattern is often useful for specialized notations and domain-specific languages (DSLs).

![Class Diagram](./assets/class-diagram.drawio.png)

## Why is the "Interpreter Pattern" needed?

1. **Language Representation**
If you have a domain-specific language or expressions that can be described with a simple grammar, the Interpreter Pattern provides a clean way to parse and evaluate those expressions.

2. **Ease of Extension**
You can add or modify grammar rules without breaking the entire interpreter. New interpretations often mean just implementing new expression classes.

3. **Readable and Maintainable**
The pattern decomposes complex language rules into smaller classes, improving clarity. Each class represents a single grammar component, making the overall system easier to maintain.

4. **Flexibility**
By combining multiple smaller expression objects (e.g., “TerminalExpression,” “AndExpression,” “OrExpression,” etc.), you can build sophisticated logic while keeping the core building blocks simple.

## Sample Program

In this project, you’ll see a simplified example of the Interpreter Pattern for checking whether certain keywords appear in a string:

- **Expression**(Interface): Declares the Interpret(context string) bool method that each concrete expression implements.

- **TerminalExpression**: Checks if a specific word is contained in the given context string.

- **AndExpression & OrExpression**(Composite Expressions): Evaluate multiple expressions in combination.
  - `AndExpression` returns true only if both sub-expressions are true.
  - `OrExpression` returns true if at least one of the sub-expressions is true.

- **Client Code**(in main.go): Creates and combines expression objects (e.g., “milk and sugar” or “milk or sugar”), then interprets them against various text strings. The result indicates whether both or either keyword is found in each string.

To run the sample code, execute 
```bash
go run ./cmd/interpreter/main.go
```