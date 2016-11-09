package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	suffixes := []string{".txt", ".php"}
	files := []string{"a.txt", "b.mp4", "c.php"}
	sink(filterSuffixes(suffixes, source(files)))
}

func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, file := range files {
			time.Sleep(time.Second)
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

func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}
