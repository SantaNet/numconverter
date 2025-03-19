# NumConvert

NumConvert is a utility for converting numbers between different numeral systems. It supports the following formats:

- Binary (0b)
- Octal (0o)
- Hexadecimal (0x)
- Decimal

## Features

- Automatic detection of the numeral system from input
- Conversion of numbers between different numeral systems
- Support for prefixed numbers (0x, 0b, 0o) or plain numbers
- Interactive input with command support

## Installation and Running

### Clone the repository

```sh
$ git clone https://github.com/yourusername/NumConvert.git
$ cd NumConvert
```

### Build

```sh
$ go build -o numconvert
```

### Run

```sh
$ ./numconvert
```

## Usage

After launching the program, the user can input a number with a prefix (0x, 0b, 0o) or without one. Additionally, a target numeral system can be specified for conversion.

Example inputs:

```
> 0x1A
26

> 0b1010
10

> 0o777
511

> 0x1A to bin
11010

> 777 to hex
309
```

To exit the program, type:

```
> stop
```

## Project Structure

```
NumConvert/
│── internal/
│   ├── format.go   # Format handling logic
│── main.go         # Main program file
│── go.mod          # Go module
│── go.sum          # Dependencies
```

## Contribution

If you have ideas for improvements, please create a Pull Request or open an Issue.