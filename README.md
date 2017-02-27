## Rotation logfile

This rotates the logfile for libraries that have the switching function:
```go
func SetOutput(writer io.Writer)
```

```go
r, err := logrotation.New("app.log", 0666, log.SetOutput)
...
r.Rotate()

```

Example of a rotation with the standard Logger:

```go
package main

import (
	"fmt"
	"github.com/y4v8/logrotation"
	"log"
	"os"
)

func main() {
	name := "std.log"
	r, err := logrotation.New(name, 0666, log.SetOutput)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for i := 1; i < 100; i++ {
		log.Printf("%03d", i)
		if i%10 == 0 {
			rotFileName := fmt.Sprintf("%03d.%s", i/10, name)
			os.Rename(name, rotFileName)

			r.Rotate()
		}
	}
}
```  
