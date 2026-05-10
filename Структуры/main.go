package main

import (
	"fmt"
)

// Работа со структурами

// type User struct {
// 	Name        string
// 	Age         int
// 	PhoneNumber string
// 	IsClose     bool
// 	Rating      float64
// }

// func main() {
// 	user := User{
// 		Name:        "Alex",
// 		Age:         64,
// 		PhoneNumber: "8 (928) 444 77-16",
// 		IsClose:     true,
// 		Rating:      5.5,
// 	}
// 	fmt.Println(user)
// 	user.Name = "Vladmir"
// 	fmt.Println(user.Name)

//		user.Rating = 555
//		fmt.Println(user.Rating)
//	}
type User struct {
	Name   string
	Rating float64
}

func Greeting(u User) {
	fmt.Println("hi everyone!")
	fmt.Println("My name is ", u.Name)
	fmt.Println("My rating is ", u.Rating)
	fmt.Println("")
}

func main() {
	user := User{
		Name:   "Vasia",
		Rating: 6.0,
	}
	fmt.Println("User", user)
	Greeting(user)
}
