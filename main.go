package main

import (
	"fmt"
	"os"
	"task-auto/rename"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: task-auto <subcommand> [flags]")
		os.Exit(1)
	}
	subcommand := args[1]
	switch subcommand {
	case "rename":
		rename.ExeRename()
	case "resize":
		fmt.Print("Resizing")
	default:
		fmt.Println("Unknown subcommand: ", subcommand)
		os.Exit(1)
	} 
}