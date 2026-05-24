# testgo

A simple Go program that calculates the sum of three integers.

## Description

This is a basic Go application that prompts the user to enter three integers and then calculates and displays their sum. The program includes:
- A `calculate` function that performs the addition
- A `main` function that handles user input and output
- Unit tests for the calculate function

## Features

- Interactive input for three integers
- Sum calculation with error handling
- Unit tests to verify functionality

## Getting Started

### Prerequisites

- Go 1.26.2 or higher

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/smallscaleserver/testgo.git
   ```

2. Navigate to the project directory:
   ```bash
   cd testgo
   ```

3. Build the project:
   ```bash
   go build
   ```

### Usage

Run the program:
```bash
go run main.go
```

Or build and run:
```bash
go build
./testgo
```

When prompted, enter three integers to see their sum.

## Testing

Run the unit tests:
```bash
go test
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
