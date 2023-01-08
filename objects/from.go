package objects

// #include <objects.h>
import "C"
import "unsafe"

func ListFromSlice(elements []PyObject) PyObject {
	refs := ObjectsToRefs(elements...)
	return PyObject{C.listFromSlice(C.ulonglong(len(elements)), (*unsafe.Pointer)(&(refs[0])))}
}

func TupleFromSlice(elements []PyObject) PyObject {
	refs := ObjectsToRefs(elements...)
	return PyObject{C.tupleFromSlice(C.ulonglong(len(elements)), (*unsafe.Pointer)(&(refs[0])))}
}
