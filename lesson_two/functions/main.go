package main

import (
	"fmt"
)

// функции - это блок кода, которые можно неоднократно использовать в разным местах программы
func main() {
	// вызываем функцию 3 раза
	// если у нас есть какой-то код, который мы хотим переиспользовать в разных местах несколько раз, то
	// мы можем использовать функции для этого, вынеся кода в неё
	print()
	print()
	print()

	// в данном случае мы выводим то, что записали в качестве строки в функцию
	printMyMessage("hello")
	printMyMessage("привет")

	// можно и так:
	messageOne := "bye"
	messageTwo := "пока"
	printMyMessage(messageOne)
	printMyMessage(messageTwo)

	// если не передать в ничего в функцию, которая принимает значение, то будет ошибка:
	//printMyMessage() // ошибка

	// посмотрим на функцию fmt.Println()

	// попробуем вызвать функцию, которая что-то возвращает
	returnMyName("Bob") // при запуске мы не получаес вывод на экран
	// такие функции возвращают какое-то значение, и мы можем сделать так:
	name := returnMyName("Bob") // мы можем так делать, потому что результат выполнения функции возвращает строку
	printMyMessage(name)

	// ну или так:
	var nameTwo string
	nameTwo = returnMyName("Duck")
	printMyMessage(nameTwo)
}

func print() {
	fmt.Println("функция printFunc()")
}

// функция может принимать значения
func printMyMessage(message string) {
	fmt.Println(message)
}

// функция может возвращать значение
func returnMyName(name string) string {
	// ключевое слово return, которое возвращает значение того типа, который мы определили в сигнатуре функции
	return "Имя: " + name + ";"
}

// в пакете fmt есть функция Sprintf(). посмотрим на неё. эта функция возвращает строку в некотором формате
func fmtSprintfExample(name string, balance int) string {
	// по результату выполнения функции у нас возвращается форматированная нами строка
	// эта строка подставляется после return и возвращается
	return fmt.Sprintf("Привет, %s. Твой баланс = %d", name, balance)
}
