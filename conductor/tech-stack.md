# Tech Stack: ebitenpad

## Core Technologies
- **Language:** [Go](https://go.dev/) (v1.25.0) - High-performance, statically typed language ideal for game engine libraries.
- **Framework:** [Ebitengine](https://ebitengine.org/) (v2.x) - A dead simple 2D game engine for Go.

## Architecture
- **Package Management:** Go Modules (`go.mod`).
- **Internal Structure:** Modular sub-packages for `input`, `virtual`, and `examples`.
- **API Model:** Polling-based state management that aligns with Ebitengine's `Update`/`Draw` loop.

## Platform Support
- **Desktop:** Windows, macOS, Linux.
- **Web:** WebAssembly (WASM).
- **Mobile:** Android, iOS.

## Testing & Tooling
- **Test Framework:** Standard Go `testing` package.
- **Build System:** Standard `go build` and `go test` tools.
- **Code Style:** Standard `gofmt` and `go vet`.
