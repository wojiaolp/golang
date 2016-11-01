package common

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func CommandLineFiles(files []string) []string {
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

func MinimumInt(first int, rest ...int) int {
	for _, x := range rest {
		if x < first {
			first = x
		}
	}
	return first
}

func Test() {
	fmt.Println("this is common package")
}
