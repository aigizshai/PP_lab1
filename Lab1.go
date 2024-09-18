package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//Задание 1
	TimeNow := time.Now()
	fmt.Println("Задание №1")
	fmt.Println("Текущая дата и время: ", TimeNow)
    
	//Задание 2
	fmt.Println("Задание 2")
	a := 5
	b := 5.54
	c := "Строка"
	d := true
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	//Задание 3
	fmt.Println("Задание 3")
	fmt.Println("Введите число 1")
	var first int
	fmt.Fscan(os.Stdin, &first)
	fmt.Println("Введите число 2")
	var second int
	fmt.Fscan(os.Stdin, &second)
	fmt.Println("a+b", first*second)
	fmt.Println("a-b", first-second)
	fmt.Println("a*b", first*second)
	fmt.Println("a/b", first/second)
	fmt.Println("a//b", first%second)

	//Задание 4
	fmt.Println("Введите число 1")
	var float1 float64
	fmt.Fscan(os.Stdin, &float1)
	fmt.Println("Введите число 2")
	var float2 float64
	fmt.Fscan(os.Stdin, &float2)
	float1, float2 = sum(float1, float2)
	fmt.Printf("Сумма: %f\nРазность: %f \n", float1, float2)
	fmt.Printf("Среднее значение %f\n", aver(a, first, second))
}

func sum(a, b float64) (float64, float64) {
	return a + b, a - b
}
func aver(a, b, c int) float64 {
	return ((float64(a) + float64(b) + float64(c)) / 3.0)
}
