package main

// это глобальные переменные, находящиеся вне функций
//var a int
//var b string
// важный момент: когда у нас и глобальные, и локальные переменные имеют одинаковые названия, то
// внутри функций при обращении к локальным переменным мы обращаемся именно к ним. то есть в таких случаях
// локальные переменные с таким же названием, как и у глобальных, имеют бОльший приоритет

func main() {
	// существую глобальные и локальные переменные
	//a := 3 // локальная переменная
	//b := "string"
	//fmt.Println(a, b)
}

// добавим ещё одну функцию
func printFunction() {
	//fmt.Println(a)
	// мы не можем вывести а, потому что а - это локальная переменная, находящаяся в другой функции

	// но следующий код будет корректен
	//b := "another string"
	//fmt.Println(b)
	// здесь мы создали (инициализировали) новую переменную. она является локальной для функции print()

	// сделаем так: закомментируем весь код в данной функции и попробуем напечатать в консоль a и b
	// и сравним выводы из main() и из print()
	//fmt.Println(a, b)

	// то есть если убрать инициализацию внутри функции, то мы будем использовать глобальные
}