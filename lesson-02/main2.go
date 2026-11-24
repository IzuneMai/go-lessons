package main

import "fmt"

func main() {
	const appName = "Бафузи"
	title := "Вайбик"
	cost := 500
	quantity := 10
	availability := true
	zakaz := 3

	fmt.Println("Подукт: ", title)
	fmt.Println("Заказано: ",title)
	fmt.Println("Кол-во всего: ", quantity)
	fmt.Println("Доступность: ", availability)
	fmt.Println("Итоговая сумма заказа", cost*zakaz)
}
