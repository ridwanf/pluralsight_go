package main
import "fmt"

type Salutation struct{
    name string
    greeting string
}

func Greet(salutation Salutation)  {
    fmt.Println(salutation.name);
    fmt.Println(salutation.greeting)
    fmt.Println("With function return ")
    fmt.Println(CrateMessage(salutation.name,salutation.greeting))
    fmt.Println("With function multiple return ")
    message,alternate :=CreateMessageMultiReturn(salutation.name,salutation.greeting)
    fmt.Println(message)
    fmt.Println(alternate)
}

func CrateMessage(name, greeting string) string  {
    return greeting+" "+name
}

func CreateMessageMultiReturn(name, greeting string) (string,string)  {
    return greeting+" "+name, "HEY!"+ name
}

func main()  {
    var a = Salutation{"Ridwan","Hello"}

    fmt.Println()
    Greet(a);
}
