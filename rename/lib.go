package rename

import (
	"fmt"
	"os"
)

func renameSingle(flagArgs *flagArgs) {
	if flagArgs.input == "" {
		fmt.Println("renaming a single file requires the -i (input) flag")
		os.Exit(1)
	}
	if flagArgs.output == "" {
		fmt.Println("renaming a single file requires the -o (output) flag")
		os.Exit(1)
	}

	err := os.Rename(flagArgs.input, flagArgs.output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s renamed to %s", flagArgs.input, flagArgs.output)
}