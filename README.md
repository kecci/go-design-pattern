# design-pattern-go

Install:
```sh
$ git clone https://github.com/kecci/design-pattern-go.git
```

## Creational Patterns
### Abstract Factory

> Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.

Structure:
```sh
.
├── factory
│   ├── adidas
│   │   ├── adidas.go
│   │   ├── adidasShirt.go
│   │   └── adidasShoe.go
│   ├── iSportsFactory.go
│   └── nike
│       ├── nike.go
│       ├── nikeShirt.go
│       └── nikeShoe.go
├── item
│   ├── iShirt.go
│   └── iShoe.go
└── main.go
```

## Structural Patterns
### Adapter

> Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.

Structure:
```sh
.
├── computer
│   ├── client.go
│   ├── computer.go
│   ├── mac.go
│   ├── windows.go
│   └── windowsAdapter.go
└── main.go
```

## Behavioral Patterns
### Command

> Command is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you parameterize methods with different requests, delay or queue a request’s execution, and support undoable operations.

Structure:
```sh
.
├── computer
│   ├── client.go
│   ├── computer.go
│   ├── mac.go
│   ├── windows.go
│   └── windowsAdapter.go
└── main.go
```

## Source
- https://refactoring.guru/design-patterns/go