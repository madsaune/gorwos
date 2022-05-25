# gorwos -

gorwos (Go random words or string) is a commandline tool to generate random words or strings.

## Installation

```bash
go install github.com/madsaune/gorwos@latest
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
