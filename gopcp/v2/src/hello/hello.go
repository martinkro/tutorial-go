package main

import(
    "bufio"
    "os"
    "fmt"
)

func main(){
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("Please input your name:")
    input,err := inputReader.ReadString('\n')
    if (err != nil){
        fmt.Printf("Found an error:%s\n",err)
    }else{
        fmt.Printf("len:%d\n", len(input))
        fmt.Printf("<1>%s", input)
        for i := 0; i < len(input);i++{
            fmt.Printf("%02x ", input[i])
        }
        // \r\n
        input = input[0:len(input)-2]
        fmt.Printf("<2>%s<2>", input)
        fmt.Printf("\n")
        fmt.Printf("Hello:%s!\n",input)
    }
}