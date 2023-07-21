# timezone
Is a Go-based command-line interface (CLI) app that allows users to fetch and display the current time in a specified city using its timezone.

## Installation
Needs golang 1.18+ installed.

```bash
go install github.com/usysrc/timezone@latest
```

## Usage
To use the timezone CLI, you need to provide a city name as an argument. The CLI will then fetch the current time in that city and display it in the specified 24-hour format.

Here's the basic usage of the timezone CLI:

```bash
timezone [city]
```

## Example
Let's see an example of how to use the timezone CLI. Suppose you want to know the current time in New York:

```bash
timezone America/New_York
```

Output:

```bash
Current time in New York: 15:30
```

In this example, the timezone CLI fetched the current time in New York and displayed it in the 24-hour format.

You can also specify a custom time format using the --format flag. For example, to get the time in 12-hour format with AM/PM:

```bash
timezone --format "03:04 PM" London
```

Output
```bash
Current time in London: 03:30 PM
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.