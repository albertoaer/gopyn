#include <Python.h>
#include "session.h"

PyThreadState* mainState;

void initialize() {
  // Initializes python
  Py_Initialize();
  mainState = PyThreadState_Get();
}

void finalize() {
  Py_Finalize();
}

void* mainSession(void) {
  return mainState;
}

void* newSession(void) {
  // Use always a a sub-interpreter for the session
  return Py_NewInterpreter();
}

void useSession(void* session) {
  PyThreadState_Swap(session);
}

void endSession(void* session) {
  Py_EndInterpreter(session);
}

void runString(const char* code) {
  PyRun_SimpleString(code);
}

void runFile(const char* filename) {
  PyObject *obj = Py_BuildValue("s", filename);
  FILE* file = _Py_fopen_obj(obj, "r");
  if (file)
    PyRun_SimpleFileEx(file, filename, 1);
}