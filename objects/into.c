#include <Python.h>
#include "objects.h"

const char* toRepr(void* obj) {
  PyObject* objstr = PyObject_Repr(obj);
  const char* str = PyUnicode_AsUTF8(objstr);
  Py_DECREF(objstr);
  return str;
}

const char* toStr(void* obj) {
  PyObject* objstr = PyObject_Str(obj);
  const char* str = PyUnicode_AsUTF8(objstr);
  Py_DECREF(objstr);
  return str;
}

const char* asStr(void* obj) {
  const char* str = PyUnicode_AsUTF8(obj);
  return str;
}

unsigned int asBool(void* obj, int* error) {
  int true = Py_IsTrue(obj);
  *error = true == -1;
  return true == 1;
}

long asInt(void* obj, int* error) {
  long value = PyLong_AsLong(obj);
  *error = PyErr_Occurred() != NULL;
  return value;
}

double asDouble(void* obj, int* error) {
  double value = PyFloat_AsDouble(obj);
  *error = PyErr_Occurred() != NULL;
  return value;
}