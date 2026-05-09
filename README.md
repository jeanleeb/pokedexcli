# Pokedex CLI

A simple interactive Pokedex command-line tool written in Go, built as part of [boot.dev](https://boot.dev)'s "Build a Pokedex in Go" project. It uses the [PokeAPI](https://pokeapi.co/) to fetch real Pokemon data.

The main repository is on codeberg: [https://codeberg.org/jeanleeb/pokedexcli](https://codeberg.org/jeanleeb/pokedexcli). Github repository is a mirror.

## Features

- Browse Pokemon location areas (`map`, `mapb`)
- Explore areas to find Pokemon (`explore`)
- Catch Pokemon and build your Pokedex (`catch`)
- Inspect caught Pokemon stats (`inspect`)
- View your collection (`pokedex`)
- In-memory HTTP response caching for better performance

## Installation

Requires [Go](https://go.dev/) 1.26+.

```bash
git clone https://github.com/jeanleeb/pokedexcli.git
cd pokedexcli
go build -o pokedexcli
```

## Usage

Start the REPL:

```bash
./pokedexcli
```

Alternatively, run without building with:

```bash
go run .
```

### Commands

| Command | Description |
|---------|-------------|
| `help` | Show available commands |
| `map` | List the next 20 location areas |
| `mapb` | List the previous 20 location areas |
| `explore <area>` | List Pokemon found in an area |
| `catch <pokemon>` | Try to catch a Pokemon |
| `inspect <pokemon>` | Inspect a caught Pokemon's stats |
| `pokedex` | List all caught Pokemon |
| `exit` | Quit the CLI |

## Project Structure

```
.
├── main.go                     # Entry point
├── repl.go                     # REPL loop and command registry
├── command_*.go                # CLI command implementations
├── internal/
│   ├── pokeapi/                # PokeAPI HTTP client
│   └── pokecache/              # In-memory response cache
└── go.mod
```

## License

MIT
