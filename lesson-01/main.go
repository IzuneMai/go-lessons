package main

import (
 "errors"
 "fmt"
)

//Задание 1. Информация о пользователе
// func main() {
//  var name string = "Sonya"
//  var age int = 20
//  var isDeveloper bool = true

//  fmt.Println("Имя: ",name)
//  fmt.Println("Возраст: ",age)
//  fmt.Println("Разраб: ",isDeveloper)
//  if age >= 18 {
//   fmt.Println("старость не радость")
//  }
// }

//Задание 2. Проверка доступа
// func checkAccess(age int) bool {
//  if age >= 18 {
//   return true
//  } else {
//   return false
//  }
// }

// func main() {
//  age1 := 18
//  age2 := 20
//  age3 := 10

//  fmt.Println("age1: ", checkAccess(age1))
//  fmt.Println("age2: ", checkAccess(age2))
//  fmt.Println("age3: ", checkAccess(age3))
// }

// Задание 3. Калькулятор
func calculate(a, b float64, operation string) (float64, error) {
 if operation == "+" {
  return a + b, nil
 }
 if operation == "-" {
  return a - b, nil
 }
 if operation == "*" {
  return a * b, nil
 }
 if operation == "/" {
  if b == 0 {
   return 0, errors.New("деление на 0")
  }
  return a / b, nil
 }
 return 0, errors.New("неизвестная операция")
}

func printResult(a, b float64, operation string) {
 result, err := calculate(a, b, operation)
 if err != nil {
  fmt.Println(err)
 } else {
  fmt.Println(result)
 }
}

func main() {
 printResult(2, 5, "+")
 printResult(5, 2, "-")
 printResult(2, 5, "*")
 printResult(10, 5, "/")
 printResult(10, 0, "/")
 printResult(2, 5, "$")
}