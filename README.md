# Design Patterns in Go

This project is a collection of implementations of the Gang of Four (GoF) design patterns using the Go programming language. The aim is to provide clear and concise examples of each pattern, demonstrating their use and benefits in real-world scenarios.

## Project Structure

The project is organized into several directories, each representing a different design pattern. Each directory contains the necessary Go files to implement the pattern, along with a `cmd` directory for sample programs that demonstrate the pattern in action.

- **cmd**: Contains main programs that demonstrate the usage of each design pattern.

- **internal/patterns**: Contains the implementation of the design patterns, and README.md to describe each pattern.

## How to Run the Sample Programs

To run the sample programs, navigate to the `cmd` directory of the pattern you wish to explore and execute the Go program using the `go run` command.

### Running the Iterator Pattern, for example

Navigate to the `cmd/iterator` directory and run the following command:

```bash
go run ./cmd/iterator/main.go
```

This will execute the sample program that demonstrates the Iterator Pattern with a fruit basket.

## References

This project references the following book:

- [Java言語で学ぶデザインパターン入門第3版(著:結城 浩)](https://www.amazon.co.jp/dp/B09HK66P5X)