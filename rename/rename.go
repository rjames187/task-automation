package rename

import (
	"flag"
	"fmt"
	"os"
)

func ExeRename() {
	input := flag.String("i", "/", "Defines the path of the input for the command")
	output := flag.String("o", "/", "Defines the path of the output for the command")
	newName := flag.String("new", "renamed", "Defines the new name of the target files")

	err := os.Rename(*input, fmt.Sprintf("%s%s", *output, *newName))
	if err != nil {
		fmt.Printf("Error renaming file: %s", *input)
		os.Exit(1)
	}

	fmt.Println("File renamed successfully")
}