package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	groveError    string = "grove: %s\n"
	fsnotifyError string = "grove: fsnotify: %s\n"
)

func main() {
	flag.Usage = optionHelp
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf(groveError, "try 'grove --help' for more information")
		return
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
					for _, command := range args {
						go func(command string) {
							cmd := exec.Command("sh", "-c", command)
							cmd.Stderr = os.Stderr
							cmd.Stdin = os.Stdin
							cmd.Stdout = os.Stdout

							if err := cmd.Run(); err != nil {
								fmt.Printf(groveError, err)
							}
						}(command)
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
