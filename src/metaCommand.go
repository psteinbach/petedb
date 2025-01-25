package main

import (
	"fmt"
	"os"
)

func HandleCommand(cmd string) {
	switch cmd {
	case ".exit":
		os.Exit(0)
	default:
		message := fmt.Sprintf("Unknown command: %v", cmd)
		fmt.Println(message)
	}
}
