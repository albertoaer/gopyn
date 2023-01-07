package callables

//#include <callables.h>
//#include <stdlib.h>
import "C"

import (
	"errors"
	"unsafe"

	"github.com/albertoaer/gopyn/common"
)

type Function struct {
	ref unsafe.Pointer
}

func TryWrapFunction[T ~unsafe.Pointer](ref T) (Function, error) {
	if C.isCallable(unsafe.Pointer(ref)) == 1 {
		return Function{unsafe.Pointer(ref)}, nil
	}
	return Function{nil}, errors.New("reference is not a callable object")
}

func (f Function) Call(args common.PyObjectRef, kwargs common.PyObjectRef) common.PyObjectRef {
	return common.PyObjectRef(C.call(f.ref, unsafe.Pointer(args), unsafe.Pointer(kwargs)))
}
