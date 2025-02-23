// main.go
package main

/*
#include <stdio.h>
#include <stdlib.h>
#include "hello.c"
*/
import "C"
import "unsafe"

func main() {
	name := C.CString("In the Go Program....")
	defer C.free(unsafe.Pointer(name))
	C.SayHello(name)
}
