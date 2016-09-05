package main

import "fmt"


const(
    PI =3.14
    Language="GO"
)

const(
    A=2
    B=89
    C=iota
)

func main()  {
    fmt.Println(PI)
    fmt.Println(Language)
    fmt.Println(A,B,C)
}
