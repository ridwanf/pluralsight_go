package main

import "fmt"

func main()  {
    message :="Hello go world!"

    var greeting *string = &message
    *greeting="hi"
    fmt.Println(message,*greeting)
}
