# Simple File Server in Go

This is a basic Go program that creates a simple file server using the `net/http` package. It serves files from the current directory over HTTP on port 8080 by default.

## Usage

1. Make sure you have Go installed on your system.

2. Clone this repository or download the `main.go` file.

3. Run the following command in the directory containing the `main.go` file to start the file server:

   ```bash
   go run main.go
   ```

4. The file server will start, and you can access files in the current directory by visiting `http://localhost:8080` in your web browser.

## Customizing the Port

By default, the file server listens on port 8080. If you want to use a different port, you can modify the `port` variable in the `main.go` file. For example, to run the server on port 8000, change the following line:

```go
port := ":8080"
```

to

```go
port := ":8000"
```

Then, run the program again with the updated port configuration:

```bash
go run main.go
```

## License

This code is provided under the MIT License. See the [LICENSE](LICENSE) file for details.
