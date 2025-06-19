# jo

A simple command-line tool that converts key-value arguments to JSON output.

## Installation

```bash
go install codeberg.org/usysrc/belt/jo@latest
```

## Usage

### Command-line arguments
```bash
jo key1=value1 key2=value2
```
Output:
```json
{
  "key1": "value1",
  "key2": "value2"
}
```

### From stdin
```bash
echo -e "key1=value1\nkey2=value2" | jo
```

### Combined (stdin + arguments)
```bash
echo "name=John" | jo age=30 city=Boston
```

## Examples

Simple key-value pairs:
```bash
jo name=John age=30 city=Boston
```

Nested objects using bracket notation:
```bash
jo user[name]=John user[age]=30 config[debug]=true
```

Reading from stdin:
```bash
# One key-value pair per line
cat << EOF | jo
name=John
age=30
city=Boston
EOF
```

Mixed stdin and command-line:
```bash
echo "database[host]=localhost" | jo database[port]=5432 debug=true
```

## Build

```bash
go build
```

## Test

```bash
go test ./...
```