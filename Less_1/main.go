package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

func main() {
	// score := 17
	// if score < 10 {
	// 	fmt.Println("You need a more tach ...")
	// } else if score > 15 {
	// 	fmt.Println("You are so ausome dude !")
	// } else {
	// 	fmt.Println("Greate jobe !")
	// }

	/*
		score := 51
		if score == 12 {
			fmt.Println("Дюжина!")
		} else if score == 21 {
			fmt.Println("Очко!")
		} else if score == 50 {
			fmt.Println("Полтинник!")
		} else {
			fmt.Println("Ты не попал ни в одну из категорий .")
		}
	*/
	/*
		subscribed := false
		if subscribed {
			fmt.Println("Спасибо за подписку")
		} else {
			fmt.Println("Пожалуйта , подпишись на мой канал")
		}
	*/
	// score := 12
	// if score < 6 || score > 16 {
	// 	fmt.Println("Ты попал в петушатик ! ")
	// }
	// if score > 8 && score < 15 {
	// 	fmt.Println("Ты попал в ябочко!")
	// }

	// sunny := true
	// weekend := true

	// if sunny && weekend {
	// 	fmt.Println("Я иду гулять!")
	// } else {
	// 	fmt.Println("Я сопротивляюсь")
	// }

	// computerClub := true
	// icecream := false

	// if computerClub || icecream {
	// 	fmt.Println("Я иду гулять!")
	// } else {
	// 	fmt.Println(" Я сопротивляюсь!")
	// }

	// score := 15

	// if score >= 10 {
	// 	fmt.Println("Ты красавчик ! ")
	// } else {
	// 	fmt.Println("Тбе нужно многомуу научится !")
	// }

	// score := 7
	// if score != 7 {
	// 	fmt.Println("Ты выиграл")
	// } else {
	// 	fmt.Println("ПРОЕБАЛИ ! ")
	// }

	// subscribed := false
	// fmt.Println("до условия")
	// if !subscribed {
	// 	fmt.Println("Подпишись на меня! ")
	// }

	// fmt.Println("после  условия")

	// Циклы

	// score := 0
	// for i := 1; i <= 3; i++ {

	// 	fmt.Println("Вы пролетаете трубу ", i)
	// 	fmt.Println("")
	// 	fmt.Println("Вы пролетели трубу  ", i)
	// 	score++
	// 	fmt.Println("")
	// 	fmt.Println("Ваш счет ", score)
	// 	fmt.Println("")
	// }
	//  Бесконеные циклы
	// "math/rand"
	// "time"
	// fmt.Println("Hello World !")
	// fmt.Println("random number : ", rand.Intn(10))
	// if rand.Intn(2) == 1 {
	// 	fmt.Println("Generate 1 . Congratylation ")
	// } else {
	// 	fmt.Println(" Generate any number . We miss one (1) ...")
	// }

	// score := 0

	// fmt.Println("Get ready !")
	// fmt.Println("Your score is : ", score)
	// fmt.Println("")

	// for {
	// 	fmt.Println("---------------")
	// 	fmt.Println("Вы подлетаете к трубе !")
	// 	fmt.Println("")
	// 	if rand.Intn(7) == 1 {
	// 		fmt.Println("Вы врезались в трубу :( ")
	// 		fmt.Println("Game over !")
	// 		fmt.Println("Score : ", score)
	// 		break
	// 	}
	// 	fmt.Println("---------------")
	// 	fmt.Println("Вы пролетели трубу !")
	// 	fmt.Println("")
	// 	score++
	// 	fmt.Println("Score : ", score)
	// 	fmt.Println("------------")
	// 	fmt.Println("")
	// 	time.Sleep(1000 * time.Millisecond)
	// }
	// fmt.Println("До вызова функции ")
	// fmt.Println("После вызова функции ")
	// square(5)
	// square(10)
	// square(7)
	greeting("Petr")
	number := sum2(1, 2)
	fmt.Println(number)
	fmt.Println("number", number)
}

func hello() {
	fmt.Println("Hello 1 ")
	fmt.Println("Hello 2 ")
}

func square(x int) {
	fmt.Println(" X  = ", x)
	fmt.Println(" X ^ 2 = ", x*x)
}
func sum(a int, b int) {
	s := a + b
	fmt.Println(s)
}
func greeting(name string) {
	fmt.Println("Hello dear ", name, " !")
}
func sum2(a int, b int) int {
	s := a + b
	return s
}
