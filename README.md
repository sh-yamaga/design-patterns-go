# Design Patterns in Go

This project is a collection of implementations of the Gang of Four (GoF) design patterns using the Go. The aim is to provide clear and concise examples of each pattern, demonstrating their use and benefits in real-world scenarios.

## Project Structure

The project is organized into several directories, each representing a different design pattern. Each directory contains the necessary Go files to implement the pattern, along with a `cmd` directory for sample programs that demonstrate the pattern in action.

- **cmd**: Contains main programs that demonstrate the usage of each design pattern.

- **docs**: Contains README.md to describe each pattern.

- **internal/patterns**: Contains the implementation of the design patterns.

## How to Run the Sample Programs

1. Open this project in a devcontainer to ensure all dependencies and configurations are correctly set up. 
2. Once inside the devcontainer, navigate to the `cmd` directory of the pattern you wish to explore and execute the Go program using the `go run` command.

for instace:
```bash
go run ./cmd/iterator/main.go
```

This will execute the sample program that demonstrates the Iterator Pattern with a fruit basket.

## References

This project references the following books:

- [Java言語で学ぶデザインパターン入門第3版(著: 結城 浩)](https://www.amazon.co.jp/dp/B09HK66P5X)
- [パターン指向リファクタリング入門~ソフトウエア設計を改善する27の作法(著: Joshua Kerievsky, 小黒 直樹, 村上 歴)](https://www.amazon.co.jp/dp/4822282384)
- [Go言語 100Tips ありがちなミスを把握し、実装を最適化する impress top gearシリーズ(著: Teiva Harsanyi , 柴田 芳樹)](https://www.amazon.co.jp/dp/B0CFL1DK8Q)