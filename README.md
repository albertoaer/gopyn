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
}
```

## Future features
- Full support between Python and Golang basic types
- Calls to Python functions
- Embedded Golang callbacks into Python
- Reliable thread safe concurrency model