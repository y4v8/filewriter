package main

import (
	"fmt"
	"github.com/y4v8/filewriter"
	"log"
	"os"
	"time"
)

func main() {
	const name = "std.log"

	file := filewriter.New(name, 0666)
	if err := file.Create(); err != nil {
		panic(err)
	}
	defer file.Close()

	log.SetOutput(file)

	i, n := 0, 0
	rotate := time.Tick(time.Millisecond * 100)
	tick := time.Tick(time.Millisecond * 5)
	end := time.After(time.Second * 5)
	for {
		select {
		case <-rotate:
			// external rename
			fileName := fmt.Sprintf("%03d.%s", n, name)
			os.Rename(name, fileName)

			file.Create()
			n++
		case <-tick:
		case <-end:
			return
		}
		log.Printf("%03d", i)
		i++
	}
}
