// Практика
// 1) Создай переменные товара: название, цена, количество и доступность. 
// 2) Рассчитай стоимость заказа. 
// 3) Преобразуй int в float64 и объясни, зачем это нужно. 
// 4) Создай константу с названием приложения.


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
