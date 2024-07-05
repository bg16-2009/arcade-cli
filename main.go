package main

import (
	"fmt"
	"os"
)

func main() {
	command_args := os.Args[1:]
	if len(command_args) < 1 {
		fmt.Println("no arguments specified\nTry \"arcade help\"")
		return
	}
	command := command_args[0]
	switch command {
	case "help":
	case "session":
	case "stats":
	case "goals":
	case "history":
	case "start":
	case "pause":
	case "stop":
	}
}
