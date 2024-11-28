# Nibs

Nibs is a tool for automatically bundling and running LÖVE projects. It watches for changes in your project directory and automatically creates a `.love` file and starts the LÖVE engine.

## Features

- Watches for changes in your project directory
- Automatically bundles the project into a `.love` file
- Starts the LÖVE engine with the bundled project
- Supports various file types: `.lua`, `.png`, `.jpg`, `.ogg`, `.wav`, `.frag`, `.vert`

## Requirements

- Go 1.23.2 or later
- LÖVE

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/usysrc/nibs.git
    cd nibs
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the project:
    ```sh
    go run main.go
    ```

2. Make changes to your project files (e.g., `.lua`, `.png`, `.jpg`, etc.). The tool will automatically detect changes, bundle the project, and start LÖVE2D.

## Project Structure

- `main.go`: The main Go file that contains the logic for watching files, bundling the project, and starting LÖVE2D.
- `main.lua`: The main Lua file for your LÖVE2D project.
- `Makefile`: A simple makefile to run the LÖVE2D project.
- `game.love`: The bundled LÖVE2D project file.
- `.gitignore`: Git ignore file.
- `go.mod`: Go module file.
- `go.sum`: Go dependencies file.
- `LICENSE`: License file.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.