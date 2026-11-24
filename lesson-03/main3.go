// Практика
// 1) Напиши проверку: пользователь старше 18 и имеет аккаунт. 
// 2) Добавь условие, что аккаунт не заблокирован. 
// 3) Проверь, входит ли число в диапазон от 10 до 20 включительно. 
// 4) Рассчитай остаток от деления на 2 и определи чётность.


package main

import (
 "errors"
 "fmt"
)


func accBlock(acc bool) (bool, error){
	if acc != true {
		return false, errors.New("Доступ запрещен")
	} 
	fmt.Println("Все ок")
	return true, nil
}

func main(){
	acc := true
	age := 17
	delenie := age % 2

	if age >= 10 && age <= 20{
		fmt.Println(accBlock(acc))
		fmt.Println("Остаток от деления:" , delenie)
		if delenie % 2 == 0 {
			fmt.Println("Число четное")
		} else{
			fmt.Println("Число нечетное")
		}
	} else {
		fmt.Println("Тебе сюда нельзя")
	}
}