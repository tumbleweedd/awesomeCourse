package main

func main() {
	// некоторые условные операторы
	//var a int = 8
	//var b int = 3
	//var c bool = a == b
	//fmt.Println(c)      // false

	//var a int = 8
	//var b int = 3
	//var c bool = a > b   // true

	//var a int = 8
	//var b int = 3
	//var c bool = a <= b  // false

	isValidBalance(15)
	isValidBalance(-1)
}

// множественное возвращаемое значение в функции и if_else
func isValidBalance(balance int) (string, bool) {
	// условный оператор. пишем if, затем условие
	if balance >= 0 {
		// когда мы встречаем return, функция заканчивается
		// мы выходим из функции и присваиваем эти значения переменным
		// то есть код дальше не выполняется
		return "Вы можете пополнить счёт", true
	}
	// если условие не сработало, то мы идём далее по функции
	return "Баланс некорректный", false
}

// множественное возвращаемое значение в функции и if_else
func isValidBalanceSecondVersion(balance int) (string, bool) {
	// условный оператор. пишем if, затем условие
	if balance >= 0 {
		// когда мы встречаем return, функция заканчивается
		// мы выходим из функции и присваиваем эти значения переменным
		// то есть код дальше не выполняется
		return "Вы можете пополнить счёт", true
	} else {
		// можно через else
		return "Баланс некорректный", false
	}
	// любое условие может вернуть либо true, либо false

	//если у нас true, то выполняется один блок кода, если false, то другой блок
}
