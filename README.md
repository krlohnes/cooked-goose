# Cooked Goose 🐤

Cooked Goose is a simple command-line utility for processing SQL files in a directory. It supports filtering files by type (`*.up.sql` or `*.down.sql`), interpolating environment variables, and outputting the processed files to a new `_cooked` directory while maintaining the original directory structure.

## Features

- **Filter Files**: Process all files by default, or filter only `*.up.sql` or `*.down.sql` files with flags.
- **Environment Variable Interpolation**: Replace variables wrapped in `${}` using `github.com/mfridman/interpolate`.
- **Maintain Directory Structure**: Output processed files to a structured `_cooked` directory.
- **Error Handling**: Continues processing files even when errors occur, logging issues for debugging.
- **Overwrite Options**: Overwrites output directories, with an optional flag for confirmation.

---

## Installation

### Prerequisites
- **Go**: You need Go installed on your system. You can download it [here](https://go.dev/dl/).

### Steps to Install

1. Clone the repository:
   ```bash
   git clone https://github.com/krlohnes/cooked-goose.git
   cd cooked-goose
   ```

2. Build the executable:
   ```bash
   go build -o cooked-goose cmd/main.go
   ```

3. (Optional) Install globally:
   ```bash
   sudo mv cooked-goose /usr/local/bin/
   ```

Now, you can use `cooked-goose` from anywhere on your system.

---

Alternatively, you can do: 

```
go install github.com/krlohnes/cooked-goose@latest
```
Make sure that you have `$GOBIN` in your `$PATH`. See `https://go.dev/ref/mod#go-install` for more information

## Usage

### Command Syntax
```bash
cooked-goose [directory] [flags]
```

### Flags
| Flag            | Description                                                                 |
|------------------|-----------------------------------------------------------------------------|
| `--up`          | Process only `*.up.sql` files.                                              |
| `--down`        | Process only `*.down.sql` files.                                            |
| `--overwrite`   | Overwrite the output `_cooked` directory if it already exists.              |
| `--output-dir`  | Specify a custom output directory (default is a sibling to the passed directory called `[directory]_cooked`).        |
| `--help`        | Display help information for the command.                                   |

### Examples

#### Process All SQL Files
Process all `.sql` files in the `migrations` directory:
```bash
cooked-goose migrations
```

#### Filter `*.up.sql` Files
Process only `*.up.sql` files in the `migrations` directory:
```bash
cooked-goose migrations --up
```

#### Filter `*.down.sql` Files
Process only `*.down.sql` files in the `migrations` directory:
```bash
cooked-goose migrations --down
```

#### Output to a Custom Directory
Output processed files to the `custom_cooked` directory:
```bash
cooked-goose migrations --output-dir custom_cooked
```

#### Overwrite the `_cooked` Directory
Process files and overwrite the existing `_cooked` directory:
```bash
cooked-goose migrations --overwrite
```

#### Display Help Information
Display detailed help information:
```bash
cooked-goose --help
```

---

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

---

## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for more information.
