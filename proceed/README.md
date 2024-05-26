# Proceed

`proceed` is a simple comamnd that is designed to ask the user for confirmation.

It prompts the user with a "y/N" confirmation message, and if the user confirms with "y", it returns an exit code of 0, indicating success. Otherwise, it returns an exit code of 1, indicating that the command execution was aborted.

## Installation

To install the `proceed` CLI program using `go install`, follow these steps:

1. Ensure you have Go installed on your machine. If not, download and install it from the official [Go website](https://golang.org/dl/).

2. Open a terminal and run the following command to install the `proceed` binary:
   ```sh
   go install github.com/usysrc/proceed
   ```

This will install the `proceed` binary into your Go bin directory, which should be included in your system's PATH environment variable.

## Usage

### Interactive Usage

You can use `proceed` interactively from the command line. Run the `proceed` command together with the command you are trying have a prompt for. For example:

```bash
proceed && rm file.txt || echo "Not proceeding or error."
```

In this example, if the user confirms the execution by typing "y", the `rm file.txt` command will be executed. Otherwise, a message indicating that the command execution was aborted will be printed.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request with improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
