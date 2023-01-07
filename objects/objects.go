package objects

//#include "objects.h"
//#include "stdlib.h"
import "C"
import (
	"unsafe"

	"github.com/albertoaer/gopyn/common"
)

type PyObject struct {
	ref unsafe.Pointer
}

func Wrap[T ~unsafe.Pointer](ref T) PyObject {
	return PyObject{unsafe.Pointer(ref)}
}

func (p PyObject) Ref() common.PyObjectRef {
	return common.PyObjectRef(p.ref)
}

func (p PyObject) Get(name string) PyObject {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	obj := C.getAttr(p.ref, cstr)
	return PyObject{obj}
}
