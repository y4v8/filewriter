package main

import (
	"github.com/y4v8/logrotation"
	"log"
	"os"
	"fmt"
	"sync"
	"time"
	"path/filepath"
	"strings"
)

const (
	logFileName  = "std.log"
	loggerNumber = 8
	testDuration = time.Second
)

func init() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, logFileName) {
			return os.Remove(path)
		}
		return nil
	})
}

func main() {
	r, err := logrotation.New(logFileName, 0666, log.SetOutput)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	quit := make(chan bool)
	timeout := time.After(testDuration)
	rotateSignal := make(chan bool)
	wg := &sync.WaitGroup{}

	for n := 0; n < loggerNumber; n++ {
		wg.Add(1)
		go parallelWrite(n, quit, wg)
	}

	go renaming(logFileName, rotateSignal, testDuration / time.Duration(10))

loop:
	for {
		select {
		case <-rotateSignal:
			r.Rotate()
		case <-timeout:
			close(quit)
			break loop
		}
	}

	wg.Wait()
}

func renaming(name string, rotateSignal chan bool, renameDelay time.Duration) {
	i := 0
	for {
		time.Sleep(renameDelay)

		rotFileName := fmt.Sprintf("%03d.%s", i, name)
		err := os.Rename(name, rotFileName)
		if err != nil {
			panic(err)
		}

		rotateSignal <- true

		i++
	}
}

func parallelWrite(n int, quit chan bool, wg *sync.WaitGroup) {
	i := 0
loop:
	for {
		select {
		case <-quit:
			wg.Done()
			break loop
		default:
			log.Printf("[%06d] test_%06d\n", n, i)
			i++
		}
	}
}
