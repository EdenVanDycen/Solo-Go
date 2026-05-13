// package main

// import (
// 	"fmt"
// )

// func main() {
// 	firstNumber := 11
// 	secondNumber := 9
// 	square(5)
// 	square(8)
// 	square(10)
// 	userName := "Alex"
// 	greeting(userName)
// 	sum := sum(firstNumber, secondNumber)
// 	fmt.Println(sum)
// }
// func greeting(name string) {
// 	fmt.Println("Приветствую Вас ,", name, "!")
// 	fmt.Println(name, ",что вы бы хотели сегодня сделать? ")
// }
// func square(x int) {
// 	fmt.Println(x)
// 	x *= x
// 	fmt.Println(x)
// }
// func sum(a int, b int) int {
// 	s := a + b
// 	return s
// }

/// Два ньюанса
// 1.return в функции ,который ничего не возвращает
// 2. Глобальные переменные

package main

import (
	"fmt"
)

var number int = 5 // Это глобальная переменная
var pi float64 = 3.14

func main() {

	defer func() {
		// Специфика defer работает по stack формату . Если несколько defer ,
		// то они будут идти с конца
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()

	kvadrat := square(5)

	fmt.Println("Квадрат", kvadrat)

	greeting("")

}
func greeting(name string) {
	if name == "" {
		fmt.Println("Вы передали пустое имя")
		return
	}
	fmt.Println("Hello dear", name)

}

func square(n int) int {
	s := n * n

	return s
}
