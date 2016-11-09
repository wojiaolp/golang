package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"mylib/common"
	"os"
	"regexp"
	"runtime"
	"time"
)

type Job struct {
	filename string
	//results  chan<- Result
}

type Result struct {
	filename string
	lino     int
	line     string
}

func main() {
	//使用所有的机器核心
	//Go语言里大多数并发程序的开始处都有这一行代码,但这行代码最终将会是多余的,因为Go语言的运行时系统会变得足够聪明以自动适配它所运行的机器
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <regexp> <files>\n")
	}
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatalf("invalid regexp: %s\n", err)
	} else {
		grep(lineRx, common.CommandLineFiles(os.Args[2:]))
	}
}

var workers = 4 //runtime.NumCPU()

func grep(lineRx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, common.MinimumInt(1000, len(filenames)))
	done := make(chan struct{}, 5)
	go addJobs(jobs, filenames)
	for i := 0; i < workers; i++ {
		go doJobs(done, lineRx, jobs, results)
	}
	/*
		go awaitCompletion(done, results)
		processResults(results)
	*/
	waitAndProcessResultsBlock(done, results)
}

func addJobs(jobs chan<- Job, filenames []string) {
	for _, filename := range filenames {
		time.Sleep(time.Second)
		fmt.Println("add job")
		jobs <- Job{filename}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, lineRx *regexp.Regexp, jobs <-chan Job, results chan<- Result) {
	//fmt.Println("doJobs")
	for job := range jobs {
		fmt.Println("do job")
		job.Do(lineRx, results)
	}
	done <- struct{}{}
	//fmt.Println("done")
}

func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		//time.Sleep(time.Second)
		//fmt.Println("done")
		<-done
	}
	close(results)
}

func processResults(results <-chan Result) {
	for result := range results {
		fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
	}
}

func (job Job) Do(lineRx *regexp.Regexp, results chan<- Result) {
	file, err := os.Open(job.filename)
	if err != nil {
		log.Printf("error: %s\n", err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for lino := 1; ; lino++ {
		line, err := r.ReadBytes('\n')
		line = bytes.TrimRight(line, "\n\r")
		if lineRx.Match(line) {
			results <- Result{job.filename, lino, string(line)}
		}
		if err != nil {
			if err != io.EOF {
				log.Printf("error:%d %s\n", lino, err)
			}
			break
		}
	}

}

func waitAndProcessResultsBlock(done <-chan struct{}, results <-chan Result) {
	for working := workers; working > 0; {
		select {
		case result := <-results:
			fmt.Printf("%s:%d:%s\n", result.filename, result.lino, result.line)
		case <-done:
			working--
		default:
			fmt.Println("balabala")
		}
	}
}
