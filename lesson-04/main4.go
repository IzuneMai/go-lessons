// Практика
// 1) Напиши определение оценки по баллам. 
// 2) Через switch обработай команды меню: add, list, delete, exit. 
// 3) По номеру месяца выведи сезон. 
// 4) Сделай проверку доступа по возрасту и статусу аккаунта. ---> lesson-03


package main

import "fmt"

// Задание 1
// func main(){
// 	score := 100

// 	if score >= 90{
// 		fmt.Println("Отлично")
// 	} else if score >= 70{
// 		fmt.Println("Хорошо")
// 	} else if score >= 40{
// 		fmt.Println("Ну такое")
// 	} else{
// 		fmt.Println("Повторяй")
// 	}
// }

//Задание 2
// func main(){
// 	command := "list"

// 	switch command{
// 	case "add":
// 		fmt.Println("Добавить элемент")
// 	case "list":
// 		fmt.Println("Создать список")
// 	case "delete":
// 		fmt.Println("Удалить элемент")
// 	case "exit":
// 		fmt.Println("Выйти")
// 	default:
// 		fmt.Println("Неизвестная команда")
// 	}
// }

//Задание 3

func main(){
	month := "7"

	switch month{
	case "1", "2", "12":
		fmt.Println("Зима")
	case "3", "4", "5":
		fmt.Println("Весна")
	case "6", "7", "8":
		fmt.Println("Лето")
	case "9", "10", "11":
		fmt.Println("Осень")
	default:
		fmt.Println("Такого месяца нет")
	}
}