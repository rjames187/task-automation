package rename

import (
	"flag"
	"fmt"
	"os"
)

type flagArgs struct {
	input string
	output string
	mult bool
	inDir string
	outDir string
}

func newFlagSet() (*flag.FlagSet, *flagArgs) {
	flags := flag.NewFlagSet("rename", flag.ContinueOnError)
	flagArgs := &flagArgs{}
	flags.BoolVar(&flagArgs.mult, "mult", false, "whether there are multiple files or not")
	flags.StringVar(&flagArgs.input, "i", "", "input path for renaming")
	flags.StringVar(&flagArgs.output, "o", "", "output name not including path")
	flags.StringVar(&flagArgs.inDir, "inDir", "", "input directory")
	flags.StringVar(&flagArgs.outDir, "outDir", "", "output directory")
	return flags, flagArgs
}

func ExeRename(args []string) {
	flags, flagArgs := newFlagSet()
	err := flags.Parse(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch flagArgs.mult {
	case false:
		renameSingle(flagArgs)
	}
	os.Exit(0)
}