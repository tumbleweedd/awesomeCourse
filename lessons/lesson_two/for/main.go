package main

import "fmt"

func main() {
	sum := 0
	//1) инициализации счётчика: i := 1. то есть это инициализация переменной
	//2) второе - это условие. До тех пор, пока условие возвращает true (то есть пока условие - истина), цикл будет выполняться дальше
	//3) третья часть - увеличиваем i на 1 (увеличиваем счётчик, который на следующей итерации будет проверяться на истину или ложь)
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//можно так:
	var c = 1
	for ; c < 10; c++ {
		fmt.Println(c * c)
	}

	//or this shit
	var k = 1
	for ; k < 10; {
		fmt.Println(k * k)
		k++
	}

	//or like there
	var m = 1
	for m < 10 {
		fmt.Println(m * m)
		m++
	}

	//infinity for loop
	for {

	}
}
