package main

import (
	"errors"
	"fmt"
)

func main() {
	// как с этим работаь?
	// *ошибка всегда возвращается последним аргументом*
	// мы можем проверить, вернула ли функция ошибку или нет
	result, err := isValidBalance(15)
	// если выражение - истина (то есть функцию вернула ошиьку), выполняем действие
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

// разберём рассмотренную ранее функцию
func isValidBalance(amount int) (string, error) {
	var result int
	if amount == 0 {
		result = result + amount
		return "Бедолага, зачем тебе класать 0 рублей на счёт? Что ж, ладно...", nil
	} else if amount > 0 {
		result += amount
		return "Пополнение произошло", nil
	}
	return "", errors.New("не удалось пополнить счёт")
}
