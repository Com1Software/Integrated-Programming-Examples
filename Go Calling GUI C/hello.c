// hello.c
#include <windows.h>
#include <stdio.h>

void SayHello(const char* name) {
    printf(" In the C code... %s!\n", name);
    MessageBox(NULL, name, "Message", MB_OK);
    
}