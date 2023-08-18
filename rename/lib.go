package rename

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

func renameMany(flagArgs *flagArgs) {
	if flagArgs.inDir == "" {
		fmt.Println("renaming many files requires the -inDir flag")
		os.Exit(1)
	}
	if flagArgs.outDir == "" {
		fmt.Println("renaming many files requires the -outDir flag")
		os.Exit(1)
	}
	if flagArgs.outName == "" {
		fmt.Println("renaming many files requires the -on (output name) flag")
		os.Exit(1)
	}
	i := 0
	err := filepath.Walk(flagArgs.inDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			i++
			path = flipSlashes(path)
			ext, _, dir, err := segregateFilePath(path)
			if err != nil {
				return err
			}
			newPath := fmt.Sprintf("%s%s%d.%s", dir, flagArgs.outName, i, ext)
			err = os.Rename(path, newPath)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("All files in %s renamed", flagArgs.inDir)
}