package main

import (
	"path/filepath"
	"strings"
)

func main() {
	files := []string{"a.txt", "b.mp4", "c.php"}
	rst := source(files)
	fmt.Println(rst)
}

func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, file := range files {
			out <- file
		}
		close(out)
	}()
	return out
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, len(in))
	go func() {
		for file := range in {
			if len(suffixes) == 0 {
				out <- file
				continue
			}
			ext := strings.ToLower(filepath.Ext(file))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- file
					break
				}
			}

		}
		close(out)
	}()
	return out
}
