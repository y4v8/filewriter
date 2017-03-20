package main

import (
	"context"
	"fmt"
	"github.com/y4v8/filewriter"
	"log"
	"os"
	"sync"
	"time"
)

const name = "std.log"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	file := filewriter.New(name, 0666)
	if err := file.Create(); err != nil {
		panic(err)
	}
	defer file.Close()

	log.SetOutput(file)

	rotated := make(chan bool)
	go remote(rotated)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case <-rotated:
				file.Create()
			}
		}
	}()

	for number := 0; number < 14; number++ {
		wg.Add(1)

		go func(number int) {
			defer wg.Done()

			i := 0
			tick := time.Tick(time.Millisecond * 5)
			for {
				select {
				case <-ctx.Done():
					return
				case <-tick:
				}
				log.Printf("%03d %03d", number, i)
				i++
			}
		}(number)
	}

	wg.Wait()
}

func remote(signal chan bool) {
	i := 0
	for {
		time.Sleep(time.Millisecond * 100)

		fileName := fmt.Sprintf("%03d.%s", i, name)
		err := os.Rename(name, fileName)
		if err == nil {
			signal <- true
		}
		fmt.Println("rename:", fileName, err)
		i++
	}
}
