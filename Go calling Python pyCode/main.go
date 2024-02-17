package main

// #cgo pkg-config: python3-embed
// #include "/usr/include/python3.11/Python.h"
import "C"

import (
   "unsafe"
  )
 
 func main() {

   pycodeGo := "print('hello world')\n" 
   pycodeGo = pycodeGo + "print('hello world Two')" 
   
   defer C.Py_Finalize()
   C.Py_Initialize()
   pycodeC := C.CString(pycodeGo)
   defer C.free(unsafe.Pointer(pycodeC))
   C.PyRun_SimpleString(pycodeC)
 
 }
