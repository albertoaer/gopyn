package objects

//#include "objects.h"
//#include "stdlib.h"
import "C"
import "unsafe"

type PyObject struct {
	ref unsafe.Pointer
}

func PyObjectFrom(ref unsafe.Pointer) PyObject {
	return PyObject{ref}
}

func (p PyObject) Get(name string) PyObject {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	obj := C.getAttr(p.ref, cstr)
	return PyObjectFrom(obj)
}
