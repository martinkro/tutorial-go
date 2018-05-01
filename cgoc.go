package main
 
/*
#include <stdio.h>
 
extern void GoExportedFunc();
 
void bar() {
        printf("I am bar!\n");
        GoExportedFunc();
}
*/
import "C"
 
import "fmt"
 
//export GoExportedFunc
func GoExportedFunc() {
        fmt.Println("I am a GoExportedFunc!")
}
 
func main() {
        C.bar()
}