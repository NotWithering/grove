package main

import (
	"flag"
	"fmt"
)

var (
	debounce int
)

func init() {
	flag.IntVar(&debounce, "d", 0, "")
	flag.IntVar(&debounce, "debounce", 0, "")
}

func optionHelp() {
	fmt.Println("Usage: grove [options...] <file> [command...]")
	fmt.Println("  -d, --debounce <number> Make the program wait <number> seconds between commands")
	fmt.Println("  -h, --help              Get help and usage for grove")
}
