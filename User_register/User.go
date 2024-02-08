package main

import "fmt"

type uzer struct {
	name     string
	age      int64
	password string
}

func main() {
	var name string
	fmt.Println("Введите ваше имя: ")
	fmt.Scanf("%s\n", &name)

	var b int64
	fmt.Println("Введите возраст: ")
	fmt.Scan(&b)

	var password string
	fmt.Println("Введите пароль: ")
	fmt.Scanf("%s\n", &password)

	

	fmt.Scan(&b)
	fmt.Println("Регистрация завершена! Спасибо, за выбор наших услуг!")
	fmt.Scan(&b)
}


