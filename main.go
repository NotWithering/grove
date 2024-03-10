package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	groveError    string = "grove: %s\n"
	fsnotifyError string = "grove: fsnotify: %s\n"
)

var (
	args []string

	debounce int = 0
)

func main() {
	args = os.Args
	args = args[1:] // trim $0

	if len(args) == 0 {
		fmt.Printf(groveError, "try 'grove --help' for more information")
		return
	}

	// parsing [options...]
	for _, arg := range args {
		if arg, found := strings.CutPrefix(arg, "-"); found {
			// ex. --help
			if option, found := strings.CutPrefix(arg, "-"); found {
				switch option {
				case "help":
					optionHelp()
					return
				case "debounce":
					if optionDebounce() {
						return
					}
				}
			} else {
				// ex. -h
				for _, option := range arg {
					switch option {
					case 'h':
						optionHelp()
						return
					case 'd':
						if optionDebounce() {
							return
						}
					}
				}
			}
			args = args[1:] // trim [options...]
		} else {
			break
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf(fsnotifyError, err)
		return
	}
	defer watcher.Close()

	if len(args) == 0 {
		fmt.Printf(groveError, "no file specified")
		return
	}

	err = watcher.Add(args[0])
	if err != nil {
		fmt.Printf(fsnotifyError, err)
		return
	}
	args = args[1:]

	var debounceTime time.Time = time.Unix(0, 0)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				if debounceTime.UnixNano() > time.Now().UnixNano() {
					continue
				}
				debounceTime = time.Now().Add(time.Duration(debounce) * time.Second)

				if len(args) == 0 {
					continue
				}
				go func() {
					cmd := exec.Command(args[0], args[1:]...)
					cmd.Stderr = os.Stderr
					cmd.Stdin = os.Stdin
					cmd.Stdout = os.Stdout

					if err := cmd.Run(); err != nil {
						fmt.Printf(groveError, err)
					}
				}()
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Printf(fsnotifyError, err)
		}
	}
}
