package main

// #include <hello.c>
// void SayHello(const char* s);
/*
#include <stdio.h>
void SayBye(const char* s){
	puts(s);
}

*/
import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
	C.SayBye(C.CString("Bye, World\n"))
}
