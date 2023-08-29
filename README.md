# Texc

Cli tool that unescapes text

## Motivation

Stuctured logging is a thing and it's great. But sometimes you just want to see
the raw text. Im sure there are other tools that do this but I wanted to make
my own.

## Installation

```bash
go install github.com/karim-w/texc@latest
```

## Usage

```bash
echo "hello\nworld" | texc
```

if you use Zap logs you can pipe the output of the logs to texc

```bash
go run ./cmd 2>&1 | texc
```

> Note: the `2>&1` is to redirect the stderr to stdout so that it can be piped

## License

BSD-3-Clause

## Author

karim-w

## Contributing

PRs
