# UUID 
A simple UUID generator

## Installation

Needs golang 1.18 installed.

```bash
git clone https://github.com/usysrc/uuid
cd uuid
go install
```


## Usage

### Generate a UUID

```bash
uuid
```

```
d101c929-2d03-420a-bbdc-0a3b4549ef7f
```

### Generate multiple UUIDs

```bash
uuid -n 5
```

```
c6574032-1b0b-4c69-a9b3-4743698259da
f1868562-6fc5-409c-9232-8c377568f79b
446d376a-2eda-4648-926a-2eab0d7f5a3a
a1ac1975-3732-4c5d-8aae-825345e84f15
2f962017-1e37-408a-9076-1800bae5a0c6
```

### Verify a UUID

```bash
uuid -v c6574032-1b0b-4c69-a9b3-4743698259da
```

```
valid
```
