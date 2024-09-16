package main

import (
    "fmt"
    "time"
)    

func main() {
    //Задание 1
    TimeNow := time.Now()
    fmt.Println("Задание №1")
    fmt.Println("Текущая дата и время: ",TimeNow)
    
    //Задание 2
    fmt.Println("Задание 2")
    var int a = 5
    var float64 b = 5.54
    var string c = "Строка"
    var bool d = true
    fmt.Println("a=",a)
    fmt.Println("b=",b)
    fmt.Println("c=",c)
    fmt.Println("d=",d)

    

}

func sum(a,b:int) {
    return a+b,a-b
}
