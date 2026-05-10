package main

// Работа с  Указателями

import (
	"fmt"
)

// func main() {
// 	number := 5
// 	// Знак амперсанта (&) является указатеем и указывает на переменн number
// 	pointer := &number
// 	foo(pointer)
// 	boo(number)
// }

// // У указателей надо опредеять тип самого указателя ,
// // чтобы не ломать логику . То есть
// // pointer := &number (& указатель) .... func foo (n *int) {} (*int//
// // * это указатель указателя и ему мы определяем тип указателя . В примере *int.
// // Но определять это надо в зависимотсти от типа данных в переменной. )

// func foo(n *int /* Объявляет тип указателя*/) {
// 	fmt.Println(n)  // Выдаст адресс указателя
// 	fmt.Println(*n) // Выдаст переменную ,которая лежит в указателе
// }

//	func boo(n int) {
//		n = 10
//	}

// func main() {
// 	name := "Игорь" //Создание переменной

// 	ptr := &name // Указатель на перемнную name // name := "Игорь"

// 	fmt.Println("Имя до изменения : ", name) // Вызываем name

// 	changeName(ptr) // Вызываем функцию ,которая принимает в себя указатель
// 	// , следовательно функция КОПИРУЕТ УКАЗАТЕЛЬ . Но он указывает на ТУ ЖЕ САМУЮ ОБЛАСТЬ ПАМЯТИ . ЗАТЕМ ЛОГИКА ПЕРЕДАЕТСЯ В функцию changeName(s *string) -------

// 	fmt.Println("Имя после изменения : ", name)
// }

// func changeName(s *string) {
// 	*s = "Иван"
// 	// ----- ЗАТЕМ ЛОГИКА ПЕРЕДАЕТСЯ В функцию changeName(s *string) ,где этот указатель уже разыменовывается . Затем доходит до переменной и записывает туда  	*s = "Иван" .
// 	// То есть при помощи разыменнования мы перезаписали в основную переменную "Иван" , не взаимодействуя с самой переменной напрямую
// }

// Nil указатели
// func main() {
// 	number := 15
// 	ptr := &number
// }

// Практика от нейронки
//  1
// func main() {
// 	var count int = 42
// 	ptr := &count
// 	pointer(ptr)
// 	fmt.Println(ptr)
// 	defer func() {
// 		*ptr = 100
// 		fmt.Println(count)
// 		fmt.Println(*ptr)
// 	}()
// }

// func pointer(n *int) {
// 	fmt.Println(n)
// 	fmt.Println(*n)
// }

// func main() {
// 	// foo := 23
// 	// fmt.Println(foo)
// 	// ptrFoo := &foo
// 	// fmt.Println(ptrFoo)

// 	// fmt.Println(*ptrFoo)
// 	// *ptrFoo = 2222
// 	// fmt.Println(*ptrFoo)
// 	num := 4
// 	squareVal(num)
// 	fmt.Println(&num)
// 	squareAdd(&num)
// }

// func squareAdd(p *int) {
// 	println(p)
// 	sqr := *p * *p
// 	*p = sqr
// }

// func squareVal(n int) {
// 	fmt.Println(&n)
// }

// Попытка 3 усвоить указатели

// func main() {
// 	number := 15
// 	fmt.Println(&number)
// 	fmt.Println(number)
// 	updateValue(&number)
// 	fmt.Println(&number)
// 	fmt.Println(number)

// }

// func updateValue(n *int) {
// 	*n = 12
// 	fmt.Println(&n)
// 	fmt.Println(*n)
// }

// func main() {
// 	num1 := 10
// 	num2 := 5
// 	fmt.Println(num1, num2)
// 	swap2(&num1, &num2)

// }

// //	func swap(a, b *int) {
// //		temp := *a
// //		*a = *b
// //		*b = temp
// //		fmt.Println(*a, *b)
// //	}
// func swap2(a, b *int) {
// 	*a, *b = *b, *a
// 	fmt.Println(*a, *b)
// }

func main() {
	// number := 15

	var ptr *int
	if ptr != nil {
		fmt.Println("Указатель на ", ptr)
	} else {
		fmt.Println("Указывает на nil")
	}
}
