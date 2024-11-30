# nibs

`nibs` is a tool for automatically bundling and running LÖVE projects. It watches for changes in your project directory and automatically creates a `.love` file and starts the LÖVE engine.

## Features

- Watches for changes in your project directory
- Automatically bundles the project into a `.love` file
- Starts the LÖVE engine with the bundled project
- Supports various file types: `.lua`, `.png`, `.jpg`, `.ogg`, `.wav`, `.frag`, `.vert`

## Requirements

- Go 1.23.2 or later
- LÖVE

## Installation

1. Install with go: 
```sh
go install github.com/usysrc/nibs
```

## Usage

### Add libraries
> [!warning]
> *currently only supporting hump*

Go to your LÖVE project directory and run:

```shell
nibs add hump
```

### Bundle
Go to your LÖVE project directory and run:

```sh
nibs bundle -o output.love
```

If you don't provide a `-o` option the output will be `[directory].love`.

### Watch
Go to your LÖVE project directory and run:

1. Run the project:
```sh
nibs watch
```

2. Make changes to your project files (e.g., `.lua`, `.png`, `.jpg`, etc.). The tool will automatically detect changes, bundle the project, and restart LÖVE.


## Known issues
- Focus stealing: when restarting LÖVE, the focus will shift to the newly created instance, annoying if you are in the habit of saving often.

## Project Structure

- `main.go`: The main Go file that contains the logic for watching files, bundling the project, and starting LÖVE.
- `main.lua`: The main Lua file for your LÖVE project.
- `Makefile`: A simple makefile to run the LÖVE project.
- `game.love`: The bundled LÖVE project file.
- `.gitignore`: Git ignore file.
- `go.mod`: Go module file.
- `go.sum`: Go dependencies file.
- `LICENSE`: License file.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.