#include <Python.h>
#include "objects.h"

void* listFromSlice(size_t sz, void** items) {
  PyObject* list = PyList_New(sz);
  for (size_t i = 0; i < sz; i++)
    PyList_SetItem(list, i, items[i]);
  return list;
}

void* tupleFromSlice(size_t sz, void** items) {
  return PyList_AsTuple(listFromSlice(sz, items));
}