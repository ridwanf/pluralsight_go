package main

import "fmt"

func main()  {
    message :="Hello go world!"

    var greeting *string = &message
    *greeting+="lalal"
    fmt.Println(message,*greeting)
}
