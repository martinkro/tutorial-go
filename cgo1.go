package main

/*
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

import "fmt"

func Random()int{
    return int(C.rand())
}

func Seed(i int){
    C.srand(C.uint(i))
}

func main(){
    Seed(100)
    fmt.Println("Random:",Random())
}