# obs
A CLI for interacting with Obsidian.

## Installation
Needs golang 1.23+ installed.

To use obs, you need to have Go installed on your system. If you haven't installed Go, you can download it from the official Go website: https://golang.org/dl/

Once you have Go installed, you can install obs using the following command:

```bash
go install codeberg.org/usysrc/belt/obs
```

This will download the source code and compile the binary.

## Usage

You need to first set the default vault and the default folder for new notes.

```bash
obs config --vault "mytestvault" --targetFolder "general" --vaultPath "/absolute/path/to/mytestvault"

```

You can also set the vault path using an environment variable:

```bash
export OBS_VAULT_PATH="/absolute/path/to/mytestvault"
```

### Commands

| Command | Mode | Description |
| --- | --- | --- |
| `obs create <name> [content]` | stdio | Create a note via Obsidian URI |
| `obs search <query>` | stdio | Search your vault via Obsidian URI |
| `obs get <name>` | stdio | Print note content to stdout |
| `obs open <name>` | gui | Open a note in Obsidian app |

### Creating, reading, and opening notes

Make sure that obsidian is running!

To create new notes use:
```bash
obs create mynewnote "this is my note content"
```

To print note content to stdout:
```bash
obs get mynewnote
obs get notes/project/mynewnote
```

To open notes:
```bash
obs open mynewnote
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or create a pull request.
