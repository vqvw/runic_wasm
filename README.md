# Runic WASM

Expose functions from [the Runic parser library](https://github.com/vqvw/runic) to the JavaScript global scope using WebAssembly.

## Installation

`go get github.com/vqvw/runic_wasm`

## Build

```sh
GOOS=js GOARCH=wasm go build -o main.wasm
```
