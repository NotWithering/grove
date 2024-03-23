package main

import (
	"fmt"
	"strconv"
)

func optionHelp() {
	fmt.Println("Usage: grove [options...] <file> [command...]")
	fmt.Println(" -d, --debounce <number> Make the program wait <number> seconds between commands")
	fmt.Println(" -h, --help              Get help and usage for grove")
}

func optionDebounce() bool {
	if len(args) == 1 {
		fmt.Printf(groveError, fmt.Sprintf("option %s requires parameter", args[0]))
		return true
	}

	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf(groveError, err)
		return true
	}

	debounce = n

	args = args[1:] // value will be trimmed later
	return false
}
