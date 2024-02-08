# jenv - JSON to Environment Variables

**jenv** is a simple command-line tool written in Go that reads JSON data from stdin and outputs shell commands to set the environment variables based on the JSON key-value pairs.

## Installation
Needs golang 1.20+ installed.

To use jenv, you need to have Go installed on your system. If you haven't installed Go, you can download it from the official Go website: https://golang.org/dl/

Once you have Go installed, you can install jenv using the following command:

```bash
go install github.com/usysrc/jenv@latest
```

This will download the source code and compile the binary. You can find the binary in your `$GOPATH/bin` directory.

## Usage

You can use **jenv** by piping JSON data into it from another command or file. **jenv** will then output shell commands to set the environment variables based on the JSON key-value pairs.

### Example

```bash
echo '{"key1": "value1", "key2": "value2"}' | jenv
export key1=value1
export key2=value2
```

You can then execute the output as shell commands to set the environment variables in your current shell session.

```
eval $(cat example.json|jenv)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or create a pull request.

## Acknowledgments

This tool was inspired by similar utilities and was created to provide a simple way to convert JSON data to environment variables.
