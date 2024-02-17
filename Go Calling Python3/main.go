package main

// #cgo pkg-config: python3--embed
// #include "/usr/include/python3.11/Python.h"
import "C"
import "fmt"

func main() {
	C.Py_Initialize()
	fmt.Println(C.GoString(C.Py_GetVersion()))
	C.Py_Finalize()
}
