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
	outName string
	rgx string
}

func newFlagSet() (*flag.FlagSet, *flagArgs) {
	flags := flag.NewFlagSet("rename", flag.ContinueOnError)
	flagArgs := &flagArgs{}
	flags.StringVar(&flagArgs.input, "i", "", "input path for renaming")
	flags.StringVar(&flagArgs.output, "o", "", "output name not including path")
	flags.BoolVar(&flagArgs.mult, "mult", false, "whether there are multiple files or not")
	flags.StringVar(&flagArgs.inDir, "inDir", "", "input directory")
	flags.StringVar(&flagArgs.outDir, "outDir", "", "output directory")
	flags.StringVar(&flagArgs.outName, "on", "", "output file name")
	flags.StringVar(&flagArgs.rgx, "rgx", "", "regex pattern to match")
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
	case true:
		renameMany(flagArgs)
	}
	os.Exit(0)
}