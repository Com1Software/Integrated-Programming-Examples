 1 package main
 2 
// #cgo pkg-config: python3-embed
// #include "/usr/include/python3.11/Python.h"
 5 import "C"
 6
 7 import (
 8   "unsafe"
 9 )
10 
11 func main() {
12
13   pycodeGo := `
14 import sys
15 for path in sys.path:
16   print(path)
17 ` 
18   
19   defer C.Py_Finalize()
20   C.Py_Initialize()
21   pycodeC := C.CString(pycodeGo)
22   defer C.free(unsafe.Pointer(pycodeC))
23   C.PyRun_SimpleString(pycodeC)
24 
25 }
