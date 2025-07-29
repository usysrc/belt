# urlencode

A simple command-line tool to URL-encode strings.

## Usage

```
urlencode <string to encode>...
```

The tool takes one or more strings as arguments, joins them with spaces, and prints the URL-encoded result to standard output.

## Examples

### Basic Encoding

```
$ urlencode "hello world"
hello%20world
```

### Encoding Special Characters

```
$ urlencode "a/b"
a%2Fb
```

### Multiple Arguments

```
$ urlencode "hello world" "from me"
hello%20world%20from%20me
```
