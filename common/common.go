package common

import "unsafe"

type PyObjectRef unsafe.Pointer

type PyCallable interface {
	Call(args PyObjectRef, kwargs PyObjectRef) PyObjectRef
}
