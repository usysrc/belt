# slow

`slow` reads from standard input and outputs to standard output at a controlled, slow pace.

## Installation

To install `slow`, make sure you have Go installed and configured. Then, run the following command:

```bash
go install github.com/usysrc/codeberg/belt/slow@latest
```

This will install the `slow` executable in your Go binary path (e.g., `$GOPATH/bin` or `$HOME/go/bin`). Ensure this directory is in your system's `PATH` to run `slow` directly.

## Usage

`slow` reads from standard input and writes to standard output. By default, it outputs one line at a time with a small delay.

### Basic Usage

Pipe content to `slow`:

```bash
cat your_file.txt | slow
```

Or type directly into `slow` and press Enter for each line:

```bash
slow
Hello
World
^D (Ctrl+D to signal end of input)
```

### Controlling Delay

You can control the delay between lines using the `-d` or `--delay` flag, followed by a duration. Durations can be specified in various units (e.g., "500ms", "1s", "2.5s").

Example: Delay each line by 500 milliseconds:

```bash
echo -e "Line 1\nLine 2\nLine 3" | slow -d 500ms
```

### Examples

1.  **Simulate slow output of a log file:**

    ```bash
tail -f /var/log/syslog | slow
    ```

2.  **Read a file slowly, line by line:**

    ```bash
slow < my_document.txt
    ```

3.  **Combine with other commands:**

    ```bash
echo -e "Line 1\nLine 2\nLine 3" | slow
    ```


## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
