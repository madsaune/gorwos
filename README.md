# gorwos -

gorwos (Go random words or string) is a commandline tool to generate random words or strings.

## Installation

### Linux / MacOS

Download binary

```bash
# Linux
$ curl -L https://github.com/madsaune/gorwos/releases/latest/download/gorwos_linux_amd64 -o gorwos

# MacOS
$ curl -L https://github.com/madsaune/gorwos/releases/latest/download/gorwos_darwin_amd64 -o gorwos
```

Make executable

```bash
$ chmod +x gorwos
```

Install in your `$PATH`

```bash
$ sudo mv gorwos /usr/local/bin/
```

### Windows

```pwsh
$ curl https://github.com/madsaune/gorwos/releases/latest/download/gorwos_windows_amd64.exe -o gorwos.exe
```

## Usage

```bash
Usage of gorwos:
  -c int
        number of passwords (default 1)
  -l int
        length of string (default 16)
  -o string
        custom options (u, l, d, s, p, x (default "ulds")
  -v    show version
```

### Examples

```plaintext
$ gorwos -l 32
aTdYGqqT2!vxfCVa%PlDTlq-NbY8IaBZ

$ gorwos -o ul -c 5 -l 16
KGqIHLTSkBFCaPEJ
bZdyMTRJMWrZfVoE
dBmgKbQFttfbosBr
VuajJPmXblfQiNxl
KlWKATrnArBEhZWY

$ gorwos -o uls -c 3 -l 16
bAyBLDHIJx!#gYBo
LUKVn!gEw-sA-umU
ezLw*uwOZDykD=CV
```
