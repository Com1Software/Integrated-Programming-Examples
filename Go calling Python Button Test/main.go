package main

// #cgo pkg-config: python3-embed
// #include "/usr/include/python3.11/Python.h"
import "C"

import (
   "unsafe"
  )
 
 func main() {

   pycodeGo := "print('Start Button Test')\n" 
   pycodeGo = pycodeGo + "from gpiozero import Button\n" 
   pycodeGo = pycodeGo + "button = Button(2)\n" 
   pycodeGo = pycodeGo + "button.wait_for_press()\n" 
   pycodeGo = pycodeGo + "print('You pushed the Button!')\n" 
   

   defer C.Py_Finalize()
   C.Py_Initialize()
   pycodeC := C.CString(pycodeGo)
   defer C.free(unsafe.Pointer(pycodeC))
   C.PyRun_SimpleString(pycodeC)
 
 }
