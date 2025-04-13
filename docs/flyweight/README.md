# Flyweight Pattern

## What is the "Flyweight Pattern"?

The Flyweight Pattern is a structural design pattern that minimizes memory usage by sharing as much data as possible among similar objects. When you have a large number of objects that share identical or partially identical state, you can store what is common in a single shared “flyweight” and only keep unique (extrinsic) state separately.

![Class Diagram](./assets/class-diagram.drawio.png)

## Why is the "Flyweight Pattern" needed?

1. **Reduced Memory Usage**
By sharing data that’s the same across multiple objects, you avoid storing duplicate information, which can help keep your application lightweight, especially if you’re dealing with thousands or millions of similar objects.

2. **Clear Separation of Intrinsic vs. Extrinsic State**
The pattern clearly differentiates between intrinsic (shared) data stored in the flyweight, and extrinsic (unique) data that is passed in from the client or stored outside the flyweight.

3. **Efficient Object Management**
You can easily add or modify objects that share the same intrinsic data without creating and managing entirely new objects in memory.

4. **Scalability**
By reusing existing flyweights, your program can handle a much larger number of objects while keeping resource consumption under control.

## Sample Program

In this project, an example of the Flyweight Pattern is used to draw a large number of trees while sharing common internal (intrinsic) data:

- **TreeType**(Flyweight): Holds shared information, such as the tree’s name, color, and texture.

- **Tree**(Context): Represents an actual tree in the forest, containing the extrinsic state: its (x, y) coordinates. When asked to Draw(), it delegates to its TreeType to print out the shared features, plus any unique coordinate data.

- **Client Code**(in main.go): Uses TreeFactory to request a large number of trees. Trees with the same TreeType data end up sharing that flyweight. The client then calls tree.Draw() on each tree, demonstrating the shared state and unique coordinates.

To run the sample code, execute 
```bash
go run ./cmd/flyweight/main.go
```