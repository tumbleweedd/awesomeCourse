package main

import (
	"errors"
	"fmt"
)

func Control(role string) (string, int, error) {
	switch role {
	case "admin":
		return "у вас есть все права", 1, nil
	case "manager":
		return "у вас есть частичные права", 2, nil
	case "user":
		return "", 0, errors.New("forbidden")
	default:
		return "", 0, errors.New("unknown role")
	}
}

func main() {

	role := "admin"
	message, id, err := Control(role)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("%s (ID: %d)\n", message, id)
}

/*
package main

import "fmt"

func main() {
  var a float64
  fmt.Scan(&a)
  switch {
  case a > 0:
    fmt.Println(a + 10)
  case a < 0:
    fmt.Println(a - 10)
  }
}
*/
