# Gopyn (GOlang PYthon iNtegration)

Abstraction of the Python C API using CGO, developed with **Python 3.11.1**

```go
package main

import (
	"fmt"

	"github.com/albertoaer/gopyn/objects"
	"github.com/albertoaer/gopyn/session"
)

func main() {
	python := session.MainPython()
	python.RunFile("script.py")
	globals := objects.Wrap(python.Globals())
	i, e := globals.Get("some_variable").Int()
	fmt.Printf("%d %v\n", i, e) // Print the value with the possible conversion error
}
```

## Future features
- Full support between Python and Golang basic types
- Calls to Python functions
- Embedded Golang callbacks into Python
- Reliable thread safe concurrency model

## How to build with Gopyn?

It's recommended to use a building script or tool such as make, that would make easier to link every python dependency.

The following is the one I use on Windows, setting *PYTHON_PATH* as the python root directory. 
```
PYTHON_PATH=

LIBS=-lpython3 -lpython311 -l_tkinter

main:
	CGO_CFLAGS=-I$(PYTHON_PATH)/include CGO_LDFLAGS="-L$(PYTHON_PATH)/libs $(LIBS)" go build
```