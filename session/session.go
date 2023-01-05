package session

//#include "session.h"
//#include "stdlib.h"
import "C"
import (
	"unsafe"
)

type Python struct {
	session unsafe.Pointer
}

func init() {
	C.initialize()
}

func ResetAll() {
	C.finalize()
	C.initialize()
}

func MainPython() *Python {
	return &Python{C.mainSession()}
}

func NewPython() *Python {
	return &Python{C.newSession()}
}

func (p *Python) RunString(script string) {
	C.useSession(p.session)
	cstr := C.CString(script)
	defer C.free(unsafe.Pointer(cstr))
	C.runString(cstr)
}

func (p *Python) RunFile(filename string) {
	C.useSession(p.session)
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.runFile(cstr)
}
