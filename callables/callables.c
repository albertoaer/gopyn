#include <Python.h>
#include "callables.h"

int isCallable(void* obj) {
  return PyCallable_Check(obj);
}

void* call(void* obj, void* args, void* kwargs) {
  return PyObject_Call(obj, args, kwargs);
}