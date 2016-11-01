package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {

	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file> [<file2> [...<fileN>]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args

	}
	return files

}
