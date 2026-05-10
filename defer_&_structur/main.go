package main

import (
	"fmt"
)

func main() {
	fmt.Println("Я мейн и я начался")
	defer func() {
		fmt.Println("Я анонимная функция")
		hello()
		fmt.Println("Я мейн и я закончился")
	}()
	fmt.Println("Hello 1 ")
	fmt.Println("Hello 2 ")
	fmt.Println("Hello 3 ")

	hello()
}

func hello() {
	fmt.Println("Я hello и я начался")
	defer func() {
		fmt.Println("Я hello и я закончился")
	}()
	fmt.Println("Я функция hello 1")
	fmt.Println("Я функция hello 2")
	fmt.Println("Я функция hello 3")
}

// func main() {
// 	fmt.Println("Начало")
// 	defer func() {
// 		fmt.Println("Конец функции мейн")
// 	}()
// 	fool()
// }

// func fool() {
// 	fmt.Println("Start")

// 	defer func() {
// 		fmt.Println("End")
// 	}()

// 	defer func() {
// 		fmt.Println("1")
// 	}()

// 	defer func() {
// 		fmt.Println("2")
// 	}()

// 	defer func() {
// 		fmt.Println("3")
// 	}()

// }

//  Основа defer работа через stack . Все что кладется в defer будет складываться в стек и вызываться в обратном порядке  на примере ниж ,явное объяснение

/*
func main (){
 .. начало функции
 fmt.Println("Начало")
 defer func (){
 fmt.Println("Конец функции мейн")
 }()
 fool()
}

func fool (){
fmt.Println("Start fool()")

 defer func (){
 fmt.Println("End fool()")
 }()

 defer func (){
 fmt.Println("1")
 }()

 defer func (){
 fmt.Println("2")
 }()

 defer func (){
 fmt.Println("3")
 }()

}

В консоли все это будет выглядеть вот так

.."Начало"
 ........"Start"
 ........"3"
 ........"2"
 ........"1"
 ........"End"
.. "Конец функции мейн"

Все из-за особенности defer и системы stack
*/

// func main() {
// 	fmt.Println("Start main")
// 	defer func() {
// 		fmt.Println("End main")
// 	}()
// 	dataBase()
// }

// func dataBase() {
// 	// создать подключение

// 	defer func() {
// 		// закрыть стоит тут ,чтобы постоянно не следить о закрытии и учитывать
// 		// фактор ошибки человека
// 	}()

// 	boolean := true

// 	if boolean {
// 		//
// 		return
// 	} else {
// 		//
// 	}

// 	for i := 1; i <= 5; i++ {
// 		if i%2 == 0 {
// 			//
// 			return
// 		} else {
// 			//
// 			return
// 		}
// 	}
// 	// закрыть  подключение -- сноска на defer
// }
