package main

import (
	"fmt"
	"strconv"
)

func optionHelp() {
	fmt.Println("Usage: grove [options...] <file> [command...]")
	fmt.Println(" -a, --args <arguments>  Arguments to use for name option (default: -c)")
	fmt.Println(" -d, --debounce <number> Make the program wait <number> seconds between commands")
	fmt.Println(" -h, --help              Get help and usage for grove")
	fmt.Println(" -n, --name              Command to use for execution (default: bash)")
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

func optionName() bool {
	if len(args) == 1 {
		fmt.Printf(groveError, fmt.Sprintf("option %s requires parameter", args[0]))
		return true
	}
	name = args[1]

	args = args[1:]
	return false
}

func optionArgs() bool {
	if len(args) == 1 {
		fmt.Printf(groveError, fmt.Sprintf("option %s requires parameter", args[0]))
		return true
	}
	nameargs = args[1]

	args = args[1:]
	return false
}
