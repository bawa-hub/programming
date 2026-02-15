# Go Translations of Java Notes

This directory mirrors selected Java examples (SOLID and Design Patterns) in Go.

## How to run

Each example is a standalone `package main` inside its directory.

- Navigate to the example directory
- Run:

```bash
go run .
```

Examples:

- SRP (violation):
  - `cd _golang/solid/srp/violation && go run .`
- SRP (refactor):
  - `cd _golang/solid/srp/refactor && go run .`
- Strategy pattern:
  - `cd _golang/patterns/behavioral/strategy && go run .`

Go version: 1.20+ recommended. No external dependencies.
