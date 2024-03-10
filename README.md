# Grove [![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](https://github.com/NotWithering/grove/blob/main/LICENSE)

**Grove** waits for a file to be written to then runs a user specified command

## Installing
```bash
go install github.com/NotWithering/grove@latest
```

## Usage
```bash
$ grove --help
```
```
Usage: grove [options...] <file> [command...]
 -d, --debounce <number> Make the program wait <number> seconds between commands
 -h, --help              Get help and usage for grove
```

## Example
```bash
$ grove -d 1 test.txt echo hello &
$ echo "yoo" > test.txt
hello
```