package objects

//#include "objects.h"
//#include "stdlib.h"
import "C"
import "errors"

func (p PyObject) Repr() string {
	cstr := C.toRepr(p.ref)
	return C.GoString(cstr)
}

func (p PyObject) ToStr() string {
	cstr := C.toStr(p.ref)
	return C.GoString(cstr)
}

func (p PyObject) Str() (string, error) {
	cstr := C.asStr(p.ref)
	if cstr == nil {
		return "", errors.New("string conversion error")
	}
	return C.GoString(cstr), nil
}

func (p PyObject) Bool() (bool, error) {
	i := C.int(0)
	v := C.asBool(p.ref, &i)
	if i == 1 {
		return false, errors.New("boolean conversion error")
	}
	return v == 1, nil
}

func (p PyObject) Int() (int, error) {
	i := C.int(0)
	v := C.asInt(p.ref, &i)
	if i == 1 {
		return 0, errors.New("integer conversion error")
	}
	return int(v), nil
}

func (p PyObject) Float() (float64, error) {
	i := C.int(0)
	v := C.asDouble(p.ref, &i)
	if i == 1 {
		return 0, errors.New("float conversion error")
	}
	return float64(v), nil
}
