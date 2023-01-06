# Gopyn (GOlang PYthon iNtegration)

Abstraction of the Python C API using CGO

```go
package main

import (
	"fmt"

	"github.com/albertoaer/gopyn/session"
)

func main() {
	python := session.MainPython()
	python.RunFile("script.py")
	i, e := python.Globals().Get("some_variable").Int()
	fmt.Printf("%d %v\n", i, e) // Print the value with the possible conversion error
}
```

## Future features
- Full support between Python and Golang basic types
- Calls to Python functions
- Embedded Golang callbacks into Python
- Reliable thread safe concurrency model