# hex

[![Go](https://github.com/usysrc/hex/actions/workflows/go.yml/badge.svg)](https://github.com/usysrc/hex/actions/workflows/go.yml)

A simple and interactive hex viewer written in Go, leveraging the Bubble Tea library for a terminal-based UI. This tool allows you to view the hexadecimal representation of a file's contents, along with their ASCII equivalents.

## Features

- **Terminal-based UI:** Uses the Bubble Tea library for a clean, responsive UI.
- **Hex and ASCII Display:** Shows both hexadecimal and ASCII representations of file contents.
- **Interactive Scrolling:** Navigate through the file content using keyboard or mouse.
- **Dynamic Resizing:** Adjusts view to terminal window size changes.
- **High-Performance Renderer:** Option to enable high-performance rendering for complex ANSI escape sequences.

![media/hex.gif](media/hex.gif)

## Installation

Needs golang 1.21.4+ installed.

To use hex, you need to have Go installed on your system. If you haven't installed Go, you can download it from the official Go website: https://golang.org/dl/

Once you have Go installed, you can install hex using the following command:

```bash
go install github.com/usysrc/hex@latest
```

This will download the source code and compile the binary. You can find the binary in your $GOPATH/bin directory.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or create a pull request.

## Acknowledgments

This was inspired by a childhood with hex editors on DOS.
