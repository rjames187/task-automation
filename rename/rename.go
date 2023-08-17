package rename

import (
	"flag"
	"fmt"
	"os"
)

type flagArgs struct {
	input string
	output string
}

func newFlagSet() (*flag.FlagSet, *flagArgs) {
	flags := flag.NewFlagSet("rename", flag.ContinueOnError)
	flagArgs := &flagArgs{}
	flags.StringVar(&flagArgs.input, "i", "", "input path for renaming")
	flags.StringVar(&flagArgs.output, "o", "", "output path for renaming")
	return flags, flagArgs
}

func ExeRename(args []string) {
	flags, flagArgs := newFlagSet()
	err := flags.Parse(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = os.Rename(flagArgs.input, flagArgs.output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s renamed to %s", flagArgs.input, flagArgs.output)
	os.Exit(0)
}