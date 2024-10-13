# Sourced Go

Sourced Go is a robust Go implementation of event sourcing patterns, providing a solid foundation for building event-driven applications. This library empowers developers to create scalable, maintainable, and auditable systems by leveraging the power of event sourcing.

## Project Inspiration

Sourced Go is inspired by the original [sourced](https://github.com/mateodelnorte/sourced) project by Matt Walters. Patrick Lee Scott, a contributor and maintainer of the original JavaScript/TypeScript version, has brought these concepts to Go, extending and refactoring it for the Go ecosystem.

## What is Event Sourcing?

Event sourcing is an architectural pattern where the state of your application is determined by a sequence of events. Instead of storing just the current state, event sourcing systems store all changes to the application state as a sequence of events. This approach offers several benefits:

- Complete Audit Trail: Every change is recorded and can be audited.
- Temporal Query: You can determine the state of the application at any point in time.
- Event Replay: You can replay events to recreate the state or to test new business logic against historical data.

## Features

Sourced Go provides several key components to implement event sourcing in your Go applications:

- **Entity**: Represents domain objects with associated events and commands. Entities are the core of your domain model and encapsulate business logic.
- **Event Emitter**: Manages event publishing and subscription. This allows for loose coupling between components and enables reactive programming patterns.
- **Repository**: Handles entity persistence and event storage. It provides an abstraction layer over the storage mechanism, making it easy to switch between different storage solutions.
- **Event**: Represents something that has happened in the domain. Events are immutable and represent facts.

## Installation

To use Sourced Go in your project, you can install it using Go modules:

```sh
go get github.com/patrickleet/sourced_go
```

## Usage

See the [example/](https://github.com/patrickleet/sourced_go/tree/main/example) directory for sample implementations using Sourced Go. The examples demonstrate how to create entities, define events, and use the repository pattern with event sourcing.

To run the example:

```sh
go run example/main.go
```

## Project Structure

```
sourced_go/
├── pkg/
│   └── sourced/
│       ├── entity.go      # Entity interface definition
│       └── repository.go  # Repository interface and implementation
├── example/
│   ├── main.go            # Example usage
│   └── todos/             # Todo list example
│       ├── todo_repository.go
│       └── todo.go
├── go.mod                 # Go module file
└── README.md              # This file
```

## Running Tests

To run the test suite, use the following command:

```sh
go test ./...
```

## Project Status

Sourced Go is currently in active development. We are working on expanding the feature set and improving performance. Contributions and feedback are welcome!

## Roadmap

- Implement snapshotting for faster entity rebuilding
- Add support for event versioning and upcasting
- Integrate with popular databases for event storage
- Develop more comprehensive examples and documentation

## Reporting Issues

If you encounter any bugs or have feature requests, please file an issue on the [GitHub issue tracker](https://github.com/patrickleet/sourced_go/issues).

## License

This project is open-source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Thank you to all the contributors who have helped to make Sourced Go better!
