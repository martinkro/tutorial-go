// go run reflect.go
package main

import (
    "fmt"
    "reflect"
)

type T struct{
    A int
    B string
}

func ReflectStruct(){
    t := T{203,"mh203"}
    s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()
    fmt.Println("type:",typeOfT)

    for i := 0; i < s.NumField(); i++{
        f := s.Field(i)
        fmt.Printf("%d:%s %s = %v\n",i,typeOfT.Field(i).Name,f.Type(),f.Interface())
    }
}

func main(){
    var x float64 = 3.4
    fmt.Println("type:",reflect.TypeOf(x))

    v:=reflect.ValueOf(x)
    fmt.Println("type:",v.Type())
    fmt.Println("kind is float64:", v.Kind()==reflect.Float64)
    fmt.Println("value:",v.Float())

    p:=reflect.ValueOf(&x)
    fmt.Println("type of p:",p.Type())
    fmt.Println("settability of p:",p.CanSet())
    pv:=p.Elem()
    fmt.Println("settability of pv:",pv.CanSet())

    pv.SetFloat(7.1)
    fmt.Println(pv.Interface())
    fmt.Println(x)

    ReflectStruct()
}