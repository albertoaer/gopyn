#include <Python.h>
#include "objects.h"

void* getAttr(void* obj, const char* attr) {
  return PyObject_GetAttrString(obj, attr);
}