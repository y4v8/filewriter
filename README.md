## Rotation log file

Rotation log for libraries that have switching function:
```go
func SetOutput(writer io.Writer)
```

### Examples

A default standard Logger:

```go
package main

import (
	"github.com/y4v8/logrotation"
	"log"
	"os"
)

func main() {
	r, err := logrotation.New("std.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666, log.SetOutput)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for i := 1; i < 100; i++ {
		log.Printf("%03d", i)
		if i % 10 == 0 {
			rotFileName := fmt.Sprintf("std.%03d.log", i / 10)
			r.Rotate(rotFileName)
		}
	}
}
```  
  
A new standard Logger:

```go
package main

import (
	"github.com/y4v8/logrotation"
	"log"
	"os"
)

func main() {
	userLog := log.New(nil, "", log.LstdFlags)

	r, err := logrotation.New("std.user.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666, userLog.SetOutput)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for i := 1; i < 100; i++ {
		userLog.Printf("%03d", i)
		if i % 10 == 0 {
			rotFileName := fmt.Sprintf("std.user.%03d.log", i / 10)
			r.Rotate(rotFileName)
		}
	}
}

```
  
A default [logrus](https://github.com/sirupsen/logrus) Logger:

```go
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/y4v8/logrotation"
	"os"
)

func main() {
	r, err := logrotation.New("logrus.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666, log.SetOutput)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for i := 1; i < 100; i++ {
		log.Printf("%03d", i)
		if i % 10 == 0 {
			rotFileName := fmt.Sprintf("logrus.%03d.log", i / 10)
			r.Rotate(rotFileName)
		}
	}
}

```
